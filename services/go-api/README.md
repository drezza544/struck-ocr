# Go API Service

## Overview
Go API adalah service utama untuk auth, upload image, receipt management, parsing OCR, validasi, dan persistence ke PostgreSQL. Service ini berkomunikasi dengan OCR service via REST.

## Struktur Project

```
go-api/
├── cmd/                   # Entrypoint aplikasi
│   └── api/
│       └── main.go
├── config/                # Config loader
│   ├── config.go
│   └── config.yaml
├── database/              # Database setup dan migrations
│   ├── migrations/
│   │   └── 0001_init.sql
│   └── postgres/
│       ├── db.go
│       └── tx.go
├── internal/
│   ├── app/               # Bootstrap & server wiring
│   │   ├── bootstrap/
│   │   │   ├── container.go
│   │   │   └── wiring.go
│   │   └── server/
│   │       ├── http_server.go
│   │       └── middleware.go
│   ├── common/            # Shared utilities
│   │   ├── errors/
│   │   ├── logger/
│   │   ├── validator/
│   │   └── utils/
│   ├── modules/           # Feature modules (clean architecture)
│   │   ├── auth/
│   │   ├── user/
│   │   ├── upload/
│   │   ├── receipt/
│   │   ├── parsing/
│   │   ├── ocr_client/
│   │   └── health/
│   └── transport/
│       └── http/
│           ├── router.go
│           └── response.go
├── tests/                 # Unit & integration tests
│   ├── unit/
│   └── integration/
└── go.mod
```

## Penjelasan Folder dan File
- `cmd/api/main.go`: entrypoint Go API.
- `config/`: konfigurasi aplikasi dan env loader.
- `database/`: migrations dan koneksi PostgreSQL.
- `internal/app/`: wiring server dan bootstrap dependency.
- `internal/common/`: utilities lintas modul.
- `internal/modules/`: feature-based modules (auth, receipt, upload, parsing, ocr_client).
- `internal/transport/`: HTTP router dan response helpers.
- `tests/`: unit dan integration tests.

## Endpoint
- `GET /health`: health check.
