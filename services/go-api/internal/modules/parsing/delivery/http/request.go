package http

type DetectDocumentRequest struct {
	RawText string `json:"raw_text" validate:"required"`
}