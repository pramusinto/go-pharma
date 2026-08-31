# Go Pharma

Aplikasi manajemen stok obat (apotek) sederhana — dibangun sebagai project belajar full-stack dengan **React** di frontend dan **Go** di backend, terhubung ke **PostgreSQL**.

## Tech Stack

**Frontend**
- React (Vite)
- Tailwind CSS

**Backend**
- Go
- [Gin](https://github.com/gin-gonic/gin) — HTTP web framework
- [pgx](https://github.com/jackc/pgx) — PostgreSQL driver
- PostgreSQL — database

## Struktur Project

```
go-pharma/
├── backend/
│   ├── cmd/
│   │   └── api/
│   │       └── main.go          # Entry point aplikasi
│   ├── internal/
│   │   ├── config/              # Koneksi database
│   │   ├── handler/             # HTTP handler (controller)
│   │   ├── model/               # Struct data
│   │   └── repository/          # Query ke database
│   ├── go.mod
│   └── go.sum
└── frontend/
    ├── src/
    │   ├── App.jsx
    │   └── MedicineForm.jsx
    └── package.json
```

## Fitur

- [x] Menampilkan daftar obat
- [x] Menambah data obat baru
- [x] Terhubung ke database PostgreSQL
- [x] Pencarian obat berdasarkan nama/kategori
- [ ] Edit data obat
- [ ] Hapus data obat

## Cara Menjalankan

### Prasyarat

- Go 1.21+
- Node.js 20+
- PostgreSQL

### 1. Setup Database

Buat database dan tabel:

```sql
CREATE DATABASE go_pharma;

CREATE TABLE medicines (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    category VARCHAR(100) NOT NULL,
    stock INT NOT NULL DEFAULT 0,
    unit VARCHAR(50) NOT NULL,
    price NUMERIC(12,2) NOT NULL DEFAULT 0
);
```

### 2. Jalankan Backend

```bash
cd backend
```

Buat file `.env`:

```
DATABASE_URL=postgres://postgres:PASSWORD_KAMU@localhost:5432/go_pharma
```

Jalankan server:

```bash
go mod tidy
go run ./cmd/api/main.go
```

Server berjalan di `http://localhost:8080`.

### 3. Jalankan Frontend

```bash
cd frontend
npm install
npm run dev
```

Aplikasi berjalan di `http://localhost:5173`.

## API Endpoints

| Method | Endpoint              | Keterangan                          |
|--------|------------------------|--------------------------------------|
| GET    | `/api/medicines`       | Ambil semua obat (bisa `?search=`)  |
| POST   | `/api/medicines`       | Tambah obat baru                    |
| PUT    | `/api/medicines/:id`   | Update data obat                    |
| DELETE | `/api/medicines/:id`   | Hapus data obat                     |

## Roadmap

- Autentikasi user
- Riwayat transaksi/penjualan
- Deploy ke hosting gratis (frontend: Vercel/Netlify, backend: Railway/Render)