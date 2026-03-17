# OCR Receipt System

## Project Overview
OCR Receipt System adalah monorepo untuk memproses struk belanja dari gambar menjadi data receipt terstruktur yang siap disimpan dan dianalisis. Sistem terdiri dari Go API sebagai core business service dan Python OCR service untuk preprocessing + OCR extraction.

## Tujuan Sistem
- Memproses foto struk dari mobile/web client.
- Menjalankan OCR dan parsing untuk menghasilkan data receipt terstruktur.
- Menyimpan hasil akhir ke PostgreSQL.
- Menyediakan fondasi untuk analytics, categorization, dan duplicate detection.

## Tech Stack
- Backend API: Go 1.22
- OCR Service: Python 3.11 + FastAPI
- Database: PostgreSQL 16
- Containerization: Docker, Docker Compose
- Deployment target: Ubuntu

## Arsitektur Sistem (Singkat)
- Go API menerima upload, mengelola auth, receipt management, parsing OCR, dan persistence.
- OCR service menangani preprocessing image dan extraction OCR.
- Semua komunikasi antar service menggunakan REST.

Diagram detail ada di `docs/architecture.md`.

## Struktur Monorepo

```
struck-ocr/
├── services/
│   ├── go-api/         # Go API service
│   └── ocr-service/    # OCR service (FastAPI)
├── infra/              # Dockerfiles dan konfigurasi deployment
├── docs/               # Dokumentasi lengkap
├── scripts/            # Script helper dev
└── shared/             # Konfigurasi/asset lintas service
```

## Menjalankan Secara Lokal (tanpa Docker)

### Go API
```bash
cd services/go-api
GO_API_PORT=8080 go run ./cmd/api
```

### OCR Service
```bash
cd services/ocr-service
uvicorn app.main:app --host 0.0.0.0 --port 8000
```

## Menjalankan dengan Docker

```bash
make dev
```

Akses:
- Go API: `http://localhost:8080/health`
- OCR Service: `http://localhost:8000/health`

## Dokumentasi Lengkap
Semua dokumentasi detail tersedia di folder `docs/`:
- Arsitektur: `docs/architecture.md`
- OCR pipeline: `docs/ocr-pipeline.md`
- Receipt parsing: `docs/receipt-parsing.md`
- Performance & scaling: `docs/performance.md`
- Logging & monitoring: `docs/operations.md`
- Security: `docs/security.md`

## Flow Sistem End-to-End
1. User upload foto struk ke Go API.
2. Go API simpan image ke storage dan kirim URL ke OCR service.
3. OCR service preprocessing image dan menjalankan OCR.
4. OCR service mengembalikan teks mentah ke Go API.
5. Go API melakukan parsing receipt + validasi.
6. Data receipt disimpan ke PostgreSQL.
7. User dapat review/edit hasil.
