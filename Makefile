.PHONY: run-api run-ocr dev test lint migrate-up migrate-down

run-api:
	cd services/go-api && go run ./cmd/api

run-ocr:
	cd services/ocr-service && uvicorn app.main:app --host 0.0.0.0 --port 8000

dev:
	docker-compose up --build

test:
	bash scripts/test.sh

lint:
	bash scripts/lint.sh

migrate-up:
	bash scripts/migrate.sh up

migrate-down:
	bash scripts/migrate.sh down
