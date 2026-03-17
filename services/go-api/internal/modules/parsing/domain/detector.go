package domain

type Detector interface {
	Detect(rawText string) DetectionResult
}