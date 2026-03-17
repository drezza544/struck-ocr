# Security Considerations

## Input Validation
- Validasi MIME type dan ukuran file upload.
- Validasi `image_url` hanya dari domain storage internal.

## Authentication and Authorization
- JWT untuk API user.
- Restrict OCR service hanya bisa diakses dari Go API.

## Data Protection
- Enkripsi data sensitif di DB.
- Hapus image raw jika user request deletion.

## Transport Security
- Semua service komunikasi via TLS (prod).
- Gunakan signed URL untuk akses image.
