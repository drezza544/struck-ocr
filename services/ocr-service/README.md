# OCR Service (FastAPI)

## Overview
Service OCR bertanggung jawab untuk preprocessing image dan OCR extraction. Service ini stateless dan diakses oleh Go API melalui REST.

## Struktur Project

```
ocr-service/
├── app/                 # Entrypoint aplikasi FastAPI
│   ├── main.py          # Inisialisasi app + router
│   └── server.py        # Opsional: startup logic
├── config/              # Konfigurasi aplikasi
│   ├── settings.py      # Loader env/config
│   └── logging.yaml     # Konfigurasi logging
├── routers/             # HTTP routes
│   ├── health.py        # Health check endpoint
│   └── ocr.py           # OCR endpoint
├── schemas/             # Request/response schemas (Pydantic)
│   ├── request.py
│   └── response.py
├── services/            # Application/service layer
│   ├── ocr_service.py   # Orkestrasi preprocessing + OCR
│   └── parsing_service.py # Placeholder parsing di OCR service
├── ocr_engine/           # Adapter ke OCR engine
│   ├── paddle.py        # Placeholder PaddleOCR
│   └── tesseract.py     # Placeholder Tesseract
├── preprocessing/        # Modul preprocessing image
│   ├── normalize.py
│   ├── denoise.py
│   └── binarize.py
├── infrastructure/       # Integrasi infra (storage, HTTP client)
│   ├── storage/
│   │   └── local.py
│   └── http_client/
│       └── go_api.py
├── tests/                # Unit & integration tests
│   ├── unit/
│   └── integration/
└── pyproject.toml        # Dependency & project config
```

## Penjelasan Folder dan File
- `app/`: entrypoint FastAPI dan wiring router.
- `config/`: konfigurasi aplikasi (environment, logging).
- `routers/`: definisi endpoint API.
- `schemas/`: kontrak request/response.
- `services/`: business logic OCR service.
- `ocr_engine/`: adapter untuk berbagai OCR engine.
- `preprocessing/`: pipeline preprocessing.
- `infrastructure/`: koneksi ke storage/HTTP client.
- `tests/`: skenario unit/integration test.

## Endpoint
- `GET /health`: health check.
- `POST /v1/ocr`: menjalankan OCR (input: `image_url`).
