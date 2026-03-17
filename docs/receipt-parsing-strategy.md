# Receipt Parsing Strategy

## Tujuan
Mengubah teks OCR mentah menjadi data receipt terstruktur:
- `merchant_name`
- `transaction_date`
- `items`
- `subtotal`
- `tax`
- `total`
- `payment_method`

## 1) Tantangan Parsing Receipt
- Layout bervariasi antar merchant (posisi total, pajak, item).
- OCR error: karakter mirip (O/0, S/5, B/8), simbol mata uang hilang.
- Item line bisa multiline, diskon dan promo bercampur.
- Tanggal format beragam: `2024-03-01`, `01/03/2024`, `1 Mar 2024`.
- Total bisa muncul beberapa kali (subtotal, total, balance, grand total).

## 2) Deteksi Merchant Name
**Heuristik utama**
- Ambil 1–3 baris pertama yang paling “kapital” dan pendek.
- Prioritaskan baris yang mengandung kata kunci merchant (logo text) dan bukan alamat/telp.

**Aturan praktis**
- Skip baris yang mengandung `tel`, `fax`, `npwp`, `tax`, `qty`, `item`.
- Pilih baris dengan rasio huruf besar tinggi dan panjang < 30 karakter.

## 3) Deteksi Tanggal Transaksi
**Heuristik**
- Cari pola tanggal pada seluruh baris: regex untuk `YYYY-MM-DD`, `DD/MM/YYYY`, `DD-MM-YYYY`, `DD MMM YYYY`.
- Jika ada lebih dari satu tanggal, pilih yang paling dekat dengan kata kunci: `date`, `tanggal`, `trx`.

**Fallback**
- Gunakan tanggal pada header/footer yang valid (dekati jam transaksi).

## 4) Deteksi Total
**Kata kunci**
- `total`, `grand total`, `amount`, `balance`, `total bayar`.

**Heuristik**
- Ambil angka terbesar pada baris yang mengandung keyword di atas.
- Validasi: `subtotal + tax ± discount ≈ total` (toleransi rounding).

## 5) Deteksi Item Line
**Ciri item**
- Memiliki nama produk + angka harga di akhir baris.
- Pola umum: `ITEM_NAME ... PRICE`.

**Heuristik**
- Baris yang mengandung angka di akhir dan tidak berisi keyword total/tax.
- Baris setelah header `item/qty/price` sampai sebelum `subtotal/total`.

## 6) Menangani Format Receipt Berbeda
- Simpan rule per merchant (rule-based override).
- Deteksi locale/currency berdasarkan simbol (`Rp`, `$`, `RM`).
- Jika struktur tidak jelas, fallback ke rule umum.

## 7) Menangani OCR Errors
- Normalisasi karakter: `O->0`, `S->5` jika konteks numerik.
- Hapus noise berulang seperti `***`.
- Gunakan confidence threshold dari OCR engine (jika tersedia).

## 8) Validation Rules
- `total` harus >= 0.
- `subtotal + tax ± discount` mendekati `total` (toleransi 1–2%).
- Minimal 1 item atau minimal total tersedia.
- `transaction_date` harus dalam rentang valid (tidak di masa depan jauh).

## Contoh

### Input OCR Text
```
TOKO ABC
Jl. Merdeka 10
Tanggal: 01/03/2024 12:30
Item   Qty   Harga
SUSU UHT 2L    1   25.000
ROTI TAWAR     1   15.000
Subtotal           40.000
Pajak              4.000
TOTAL              44.000
Cash
```

### Output Parsing (JSON)
```json
{
  "merchant_name": "TOKO ABC",
  "transaction_date": "2024-03-01T12:30:00",
  "items": [
    {"name": "SUSU UHT 2L", "quantity": 1, "price": 25000},
    {"name": "ROTI TAWAR", "quantity": 1, "price": 15000}
  ],
  "subtotal": 40000,
  "tax": 4000,
  "total": 44000,
  "payment_method": "Cash"
}
```
