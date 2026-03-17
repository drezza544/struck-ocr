package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	ocrdomain "github.com/drezza544/struck-ocr/internal/modules/ocr_client/domain"
)

type Client struct {
	baseURL string
	client  *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

type OCRRequest struct {
	ImageURL string `json:"image_url"`
}

func (c *Client) ScanByURL(imageURL string) (*ocrdomain.OCRResponse, error) {
	reqBody := OCRRequest{
		ImageURL: imageURL,
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Post(
		c.baseURL+"/v1/ocr",
		"application/json",
		bytes.NewBuffer(bodyBytes),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("ocr service error: %s", string(body))
	}

	var result ocrdomain.OCRResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}