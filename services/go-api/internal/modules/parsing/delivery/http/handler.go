package http

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"go-api/internal/modules/parsing/service"
)

type Handler struct {
	detectorService *service.DetectorService
}

func NewHandler(detectorService *service.DetectorService) *Handler {
	return &Handler{
		detectorService: detectorService,
	}
}

func (h *Handler) DetectDocumentType(c *gin.Context) {
	var req DetectDocumentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(nethttp.StatusBadRequest, gin.H{
			"error": "invalid request body"
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