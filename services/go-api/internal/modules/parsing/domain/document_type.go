package domain

type DocumentType string

const (
	DocumentTypeUnknown				DocumentType = "unknown"
	DocumentTypeBoardingPass		DocumentType = "boarding_pass"
	DocumentTypeRetailReceipt		DocumentType = "retail_receipt"
	DocumentTypeParkingReceipt		DocumentType = "parking_receipt"
)

type DetectionResult struct {
	DocumentType	DocumentType	`json:"document_type"`
	Confidence		float64			`json:"confidence"`
	MatchedRules	[]string		`json:"matched_rules"`
}

