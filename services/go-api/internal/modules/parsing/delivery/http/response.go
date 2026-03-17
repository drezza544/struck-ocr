package http=

import ocrdomain "github.com/drezza544/struck-ocr/internal/modules/ocr_client/domain"


type DetectDocumentResponse struct {
	DocumentType	string		`json:"document_type"`
	Confidence		float64		`json:"confidence"`
	MatchedRules	[]string	`json:"matched_rules"`
}

type ProcessDocumentResponse struct {
	OCR struct {
		RawText    string                 `json:"raw_text"`
		TextBlocks []ocrdomain.OCRTextBlock `json:"text_blocks"`
	} `json:"ocr"`
	Detection DetectDocumentResponse `json:"detection"`
}