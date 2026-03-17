# Architecture Overview

## Scope
This repository contains two services and supporting infrastructure:
- Go API: auth, upload, receipt management, OCR result parsing, validation, persistence.
- OCR Service (FastAPI): image preprocessing + OCR extraction.
- PostgreSQL: persistence and queries.

## High-Level Components
- `services/go-api`: main business service.
- `services/ocr-service`: stateless OCR worker service.
- `infra/`: container build and deployment assets.
- `docs/`: documentation and runbooks.

## Backend Service Architecture

```mermaid
flowchart LR
  Client[Mobile/Web Client] -->|REST| GoAPI[Go API]
  GoAPI -->|REST| OCR[OCR Service]
  GoAPI -->|SQL| PG[(PostgreSQL)]
  GoAPI -->|Object Storage| Storage[(Image Storage)]
  OCR -->|Read Image| Storage
```

## Data Flow Antar Service

```mermaid
sequenceDiagram
  participant C as Client
  participant G as Go API
  participant S as Storage
  participant O as OCR Service
  participant P as PostgreSQL

  C->>G: Upload image
  G->>S: Store raw image
  G->>O: POST /v1/ocr (image_url)
  O->>S: Fetch image
  O->>O: Preprocess + OCR
  O-->>G: OCR text + metadata
  G->>G: Parse receipt data
  G->>P: Persist receipt
  G-->>C: Return receipt data
```

## Layering (Go API Clean Architecture)
- Domain: entities, repository interfaces, business rules.
- Application: use cases, orchestration, validation.
- Infrastructure: DB adapters, external clients.
- Transport: HTTP handlers, routing, DTO mapping.

## Where to Read More
- OCR pipeline: `docs/ocr-pipeline.md`
- Receipt parsing: `docs/receipt-parsing.md`
- Performance and scaling: `docs/performance.md`
- Logging and monitoring: `docs/operations.md`
- Security: `docs/security.md`
