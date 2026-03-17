package http

type DetectDocumentResponse struct {
	DocumentType	string		`json:"document_type"`
	Confidence		float64		`json:"confidence"`
	MatchedRules	[]string	`json:"matched_rules"`
}