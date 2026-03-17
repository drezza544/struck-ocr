package service

import (
	"strings"

	"go-api/internal/modules/parsing/domain"
)

type DetectorService struct{}

func NewDetectorService() *DetectorService {
	return &DetectorService{}
}

func (s *DetectorService) Detect(rawText string) domain.DetectionResult {
	text := normalize(rawText)

	boardingScore, boardingRules := scoreBoardingPass(text)
	retailScore, retailRules := scoreRetailReceipt(text)
	parkingScore, parkingRules := scoreParkingReceipt(text)

	bestType := domain.DocumentTypeUnknown
	baseScore := 0.0
	bestRules := []string{}

	if boardingScore > bestScore {
		bestType = domain.DocumentTypeBoardingPass
		bestScore = boardingScore
		bestRules = boardingRules
	}

	if retailScore > bestScore {
		bestType = domain.DocumentTypeRetailReceipt
		bestScore = retailScore
		bestRules = retailRules
	}

	if parkingScore > bestScore {
		bestType = domain.DocumentTypeParkingReceipt
		bestScore = parkingScore
		bestRules = parkingRules
	}

	if bestScore == 0 {
		return domain.DetectionResult{
			DocumentType: domain.DocumentTypeUnknown,
			Confidence: 0.0,
			MatchedRules: []string{},
		}
	}

	return domain.DetectionResult{
		DocumentType: bestType,
		Confidence: bestScore,
		MatchedRules: bestRules,
	}
}

func normalize(input string) string {
	text := strings.ToLower(input)
	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\t", " ")
	text = strings.Join(strings.Fields(text), " ")
	return text
}

func scoreBoardingPass(text string) (float64, []string) {
	keywords := []string{
		"boarding pass",
		"flight",
		"gate",
		"seat",
		"from",
		"to",
	}

	score := 0.0
	matched := []string{}

	for _, keyword := range keywords {
		if strings.Contains(text, keyword) {
			score += 0.16
			matched = append(matched, "contains: "+keyword)
		}
	}

	if score > 1.0 {
		score = 1.0
	}

	return score, matched
}

func scoreRetailReceipt(text string) (float64, []string) {
	keywords := []string{
		"total",
		"subtotal",
		"tunai",
		"kembalian",
		"cash",
		"change",
	}

	score := 0.0
	matched := []string{}

	for _, keyword := range keywords {
		if strings.Contains(text, keyword) {
			score += 0.16
			matched = append(matched, "contains: "+keyword)
		}
	}

	if score > 1.0 {
		score = 1.0
	}

	return score, matched
}

func scoreParkingReceipt(text string) (float64, []string) {
	keywords := []string{
		"parkir",
		"masuk",
		"keluar",
		"tarif",
		"durasi",
	}

	score := 0.0
	matched := []string{}

	for _, keyword := range keywords {
		if strings.Contains(text, keyword) {
			score += 0.16
			matched = append(matched, "contains: "+keyword)
		}
	}

	if score > 1.0 {
		score = 1.0
	}

	return score, matched
}