# Local Development

## Prerequisites
- Docker + Docker Compose
- Go 1.22
- Python 3.11

## Steps
1. Copy `.env.example` to `.env` and adjust values.
2. Run `make dev`.
3. Access Go API at `http://localhost:8080/health`.
4. Access OCR service at `http://localhost:8000/health`.
