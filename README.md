# SERVICE ACCOUNT

Sistem manajemen layanan rekening nasabah.

## 📌 FITUR
- Pendaftaran nasabah
- Sistem tabungan
- Penarikan saldo
- Pengecekan saldo
- Logging file

## 🛠 TEKNOLOGI
- **Bahasa Pemrograman**: Go/Golang
- **Framework**: Fiber
- **Database**: PostgreSQL
- **Arsitektur**: Clean Architecture

---

## 🚀 INSTALASI

### 1️⃣ Clone Repository
```bash
git clone https://github.com/nurmanhadi/service-account.git
cd service-account
go mod tidy
```

### 2️⃣ Setup Environment Variables
Buat file `.env` di dalam root proyek dan tambahkan konfigurasi berikut:
```ini
DB_USER=
DB_PASSWORD=
DB_HOST=
DB_PORT=
DB_NAME=
DB_MAX_IDLE_CONNS=
DB_MAX_OPEN_CONNS=
DB_CONN_MAX_IDLETIME= # Dalam menit
DB_CONN_MAX_LIFETIME= # Dalam menit
```

### 3️⃣ Migrasi Database
Migrasikan database di directory `pkg/database/migrations/up.sql`

### 4️⃣ Jalankan Project
```bash
go run cmd/main.go
```

Atau dengan flag `--host` dan `--port`:
```bash
go run cmd/main.go --host=127.0.0.1 --port=3000
```

### 5📌 (Opsional) Menjalankan dengan Docker
**Prerequisites**:
- Install [Docker](https://docs.docker.com/get-docker/) & [Docker Compose](https://docs.docker.com/compose/)

Jalankan perintah berikut untuk mengaktifkan layanan dengan Docker Compose:
```bash
docker-compose create
docker-compose start
docker logs service-account
```

Tunggu beberapa saat sampai log dari `service-account` muncul.
Pastikan telah mengkonfigurasi environment database di dalam file `docker-compose.yml`.

---

## 📚 DOKUMENTASI API

| Method | Endpoint | Deskripsi |
|--------|---------|-----------|
| POST   | `/nasabah/daftar` | Mendaftarkan nasabah baru |
| POST   | `/rekening/tabung` | Menambahkan saldo ke rekening |
| POST   | `/rekening/tarik` | Menarik saldo dari rekening |
| GET    | `/rekening/saldo/{no_rekening}` | Mengecek saldo rekening |

### 📌 1. POST `/nasabah/daftar`
**Request:**
```json
{
    "nama": "test",
    "nik": "34510837898",
    "no_hp": "861284221276"
}
```

**Response (200 OK):**
```json
{
    "data": {
        "no_rekening": "2001456323"
    },
    "links": {
        "self": "/api/v1/nasabah/daftar"
    }
}
```

### 📌 2. POST `/rekening/tabung`
**Request:**
```json
{
    "no_rekening": "2001407738",
    "saldo": 500000
}
```

**Response (200 OK):**
```json
{
    "data": {
        "saldo": 6000000
    },
    "links": {
        "self": "/api/v1/rekening/tabung"
    }
}
```

### 📌 3. POST `/rekening/tarik`
**Request:**
```json
{
    "no_rekening": "2001407738",
    "saldo": 500000
}
```

**Response (200 OK):**
```json
{
    "data": {
        "saldo": 5900000
    },
    "links": {
        "self": "/api/v1/rekening/tarik"
    }
}
```

### 📌 4. GET `/rekening/saldo/{no_rekening}`
**Response (200 OK):**
```json
{
    "data": {
        "saldo": 5900000
    },
    "links": {
        "self": "/api/v1/rekening/saldo/{no_rekening}"
    }
}
```

---

## 🔗 Sosial Media
[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=flat&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/nurman-hadi03)
[![GitHub](https://img.shields.io/badge/GitHub-181717?style=flat&logo=github&logoColor=white)](https://github.com/nurmanhadi)
[![Instagram](https://img.shields.io/badge/Instagram-E4405F?style=flat&logo=instagram&logoColor=white)](https://www.instagram.com/nurman00003)
[![TikTok](https://img.shields.io/badge/TikTok-000000?style=flat&logo=tiktok&logoColor=white)](https://www.tiktok.com/@manzz__3)
[![Facebook](https://img.shields.io/badge/Facebook-1877F2?style=flat&logo=facebook&logoColor=white)](https://www.facebook.com/nurmanHadi03)
[![YouTube](https://img.shields.io/badge/YouTube-FF0000?style=flat&logo=youtube&logoColor=white)](https://www.youtube.com/@nurmanhadi2457)
[![Email](https://img.shields.io/badge/Email-D14836?style=flat&logo=gmail&logoColor=white)](mailto:nurman.hadi@hotmail.com)

---

## 📜 Lisensi
Proyek ini dilisensikan di bawah [MIT License](./LICENSE). Anda bebas menggunakan, memodifikasi, dan mendistribusikan proyek ini sesuai dengan ketentuan lisensi.

