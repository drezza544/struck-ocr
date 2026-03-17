package server

import (
	"database/sql"
	"net/http"

	"struck-ocr/go-api/config"
	transport "struck-ocr/go-api/internal/transport/http"
)

func NewHTTPServer(cfg config.Config, db *sql.DB) *http.Server {
	router := transport.BuildRouter(cfg, db)
	return &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}
}
