# Performance and Scaling

## Performance Considerations
- Go API gunakan connection pooling untuk Postgres.
- Batasi ukuran upload image dan tipe file.
- Gunakan async processing untuk OCR jika traffic tinggi.
- Simpan raw OCR text untuk re-parsing tanpa OCR ulang.

## Database Indexing Strategy
- Index pada `receipts(user_id, created_at)` untuk query user history.
- Index pada `receipts(merchant_name)` jika search merchant.
- Index pada `receipt_items(receipt_id)` untuk join cepat.
- Index pada `receipts(external_hash)` untuk dedup.

## Scaling Strategy
- Scale OCR service secara horizontal (stateless).
- Go API scale horizontal di belakang load balancer.
- Postgres: gunakan read replicas untuk query berat.
- Offload OCR ke queue workers untuk volume tinggi.
