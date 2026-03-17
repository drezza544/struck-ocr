package http

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/drezza544/struck-ocr/internal/modules/parsing/service"
)

type Handler struct {
	detectorService *service.DetectorService
	ocrClient       *ocrclientservice.Client
}

func NewHandler(
	detectorService *service.DetectorService,
	ocrClient *ocrclientservice.Client,
) *Handler {
	return &Handler{
		detectorService: detectorService,
		ocrClient:       ocrClient,
	}
}

func (h *Handler) DetectDocumentType(c *gin.Context) {
	var req DetectDocumentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(nethttp.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	result := h.detectorService.Detect(req.RawText)

	resp := DetectDocumentResponse{
		DocumentType: string(result.DocumentType),
		Confidence: result.Confidence,
		MatchedRules: result.MatchedRules,
	}

	c.JSON(nethttp.StatusOK, resp)
}

func (h *Handler) ProcessDocument(c *gin.Context) {
	var req ProcessDocumentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(nethttp.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	ocrResult, err := h.ocrClient.ScanByURL(req.ImageURL)
	if err != nil {
		c.JSON(nethttp.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	detectionResult := h.detectorService.Detect(ocrResult.RawText)

	resp := ProcessDocumentResponse{
		Detection: DetectDocumentResponse{
			DocumentType: string(detectionResult.DocumentType),
			Confidence:   detectionResult.Confidence,
			MatchedRules: detectionResult.MatchedRules,
		},
	}

	resp.OCR.RawText = ocrResult.RawText
	resp.OCR.TextBlocks = ocrResult.TextBlocks

	c.JSON(nethttp.StatusOK, resp)
}