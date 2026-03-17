# Logging and Monitoring

## Logging
- Go API log structured JSON.
- OCR service log request ID dan latency.
- Semua request lintas service menyertakan correlation ID.

## Monitoring
- Metrics: request rate, error rate, latency, OCR throughput.
- Alerts: OCR error spike, DB connection saturation.
- Health checks: `/health` di Go API dan OCR service.
