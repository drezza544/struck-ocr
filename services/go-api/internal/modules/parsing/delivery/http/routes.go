package http

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, handler *Hanlder) {
	rg.POST("/document/detect", handler.DetectDocumentType)
}