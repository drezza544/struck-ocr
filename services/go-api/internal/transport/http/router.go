package http

import (
	"github.com/gin-gonic/gin"

	parsinghttp "go-api/internal/modules/parsing/delivery/http"
	parsingservice "go-api/internal/modules/parsing/service"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	
	v1 := r.Group("/v1")

	detectorService := parsingservice.NewDetectorService()
	parsingHandler := parsinghttp.NewHandler(detectorService)

	parsinghttp.RegisterRoutes(v1, parsingHandler)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return r
}