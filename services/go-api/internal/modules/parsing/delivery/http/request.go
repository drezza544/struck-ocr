package http

type DetectDocumentRequest struct {
	RawText string `json:"raw_text" validate:"required"`
}

type ProcessDocumentRequest struct {
	ImageURL string `json:"image_url" validate:"required"`
}