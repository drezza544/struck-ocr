package domain

type OCRTextBlock struct {
	Text		string	`json:"text"`
	Confidence	float64	`json:"confidence"`
}

type OCRMeta struct {
	Engine					string 	`json:"engine"`
	PreprocessingApplied	bool	`json:"preprocessing_applied"`
	OriginalFilePath		string	`json:"original_file_path"`
	ProcessedFilePath		string	`json:"processed_file_path"`
}

type OCRResponse struct {
	RawText		string			`json:"raw_text"`
	TextBlocks	[]OCRTextBlock	`json:"text_blocks"`
	Meta		OCRMeta			`json:"meta"`
}

