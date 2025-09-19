# 💳 Kampanye & Payment Service (Go + Midtrans)

Project ini adalah backend API berbasis **Golang** untuk mengelola **user, kampanye, transaksi, dan pembayaran** menggunakan **Midtrans** sebagai payment gateway.

## 🚀 Fitur Utama

- **Autentikasi & Middleware**
  - Register & login user
  - Middleware autentikasi JWT
  - Set `currentUser` berdasarkan `userID`

- **Manajemen User**
  - Registrasi & autentikasi
  - Profile user

- **Manajemen Kampanye**
  - CRUD kampanye
  - Upload gambar kampanye

- **Transaksi**
  - Membuat transaksi
  - Melihat detail transaksi kampanye
  - Integrasi dengan Midtrans untuk pembayaran

- **Payment Gateway (Midtrans)**
  - Membuat koneksi ke Midtrans
  - Proses pembayaran kampanye
  - Webhook/notification handler

## 🛠️ Teknologi yang Digunakan

- **Bahasa**: Go (Golang)
- **Framework/Library**:
  - `gorilla/mux` (routing)
  - `gorm` (ORM untuk database)
  - `midtrans-go` (payment gateway)
- **Database**: MySQL / PostgreSQL
- **Authentication**: JWT
