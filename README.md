# Go Chi API Boilerplate 🚀

[](https://golang.org/dl/)
[](https://opensource.org/licenses/MIT)
[](https://www.google.com/search?q=http://localhost:8080/swagger/index.html)

Repositori ini berisi boilerplate yang siap pakai untuk membangun REST API modern dengan Go. Proyek ini sudah dilengkapi dengan fitur-fitur penting seperti routing, autentikasi JWT, interaksi database, dokumentasi API, dan alur kerja pengembangan yang efisien menggunakan Make dan Docker.

-----

## ✨ Fitur Utama

  * **Routing Cepat**: Menggunakan [Chi](https://github.com/go-chi/chi), router yang ringan, cepat, dan idiomatis.
  * **Autentikasi JWT**: Sistem autentikasi lengkap dengan endpoint untuk registrasi, login, dan middleware untuk melindungi route.
  * **Database PostgreSQL**: Integrasi dengan PostgreSQL menggunakan driver `pgx` yang modern dan berperforma tinggi.
  * **Repository Pattern**: Struktur kode yang memisahkan logika bisnis dari logika akses data.
  * **Dokumentasi API Otomatis**: Generate dokumentasi interaktif secara otomatis dari komentar kode menggunakan [Swagger (swag)](https://github.com/swaggo/swag).
  * **Manajemen Database**: Setup database untuk development yang mudah dan konsisten menggunakan Docker Compose.
  * **Otomatisasi Tugas**: Dilengkapi `Makefile` untuk menyederhanakan perintah-perintah umum seperti menjalankan, membangun, dan migrasi database.
  * **Konfigurasi Berbasis Environment**: Membaca konfigurasi dari file `.env` untuk fleksibilitas antar lingkungan (development, production).
  * **Respon JSON Standar**: Utilitas untuk memastikan semua respon API memiliki format yang konsisten.

-----

## 🛠️ Tumpukan Teknologi (Tech Stack)

  * **Bahasa**: Go
  * **Web Framework/Router**: Chi v5
  * **Database**: PostgreSQL
  * **Driver Database**: pgx v5
  * **Autentikasi**: JSON Web Tokens (JWT)
  * **Dokumentasi**: Swagger / OpenAPI
  * **Development Environment**: Docker, Docker Compose
  * **Build/Task Runner**: Makefile

-----

## 📂 Struktur Proyek

```
/gochi-boilerplate
├── /cmd/server/
│   └── main.go           # Titik masuk aplikasi (setup server, router, db)
├── /db/
│   └── /migrations/
│       └── 001_init_schema.sql # Skema dan migrasi database
├── /docs/
│   └── ...                 # File yang di-generate oleh Swagger
├── /internal/
│   ├── /handler/           # Layer HTTP (logika request/response)
│   ├── /middleware/        # Middleware kustom (misal: autentikasi)
│   ├── /model/             # Struct untuk data (request, response, entitas)
│   ├── /repository/        # Layer akses data (interaksi dengan database)
│   └── /utils/             # Fungsi helper (JWT, respon JSON, config, dll.)
├── .env.example            # Contoh file konfigurasi environment
├── docker-compose.yml      # Konfigurasi Docker untuk database
├── go.mod                  # Manajemen dependensi Go
├── Makefile                # Shortcut untuk perintah-perintah development
└── README.md               # Dokumentasi proyek
```

-----

## 🚀 Memulai (Getting Started)

Ikuti langkah-langkah ini untuk menjalankan proyek di lingkungan lokal Anda.

### Prasyarat

Pastikan Anda sudah menginstal perangkat lunak berikut:

  * [Go](https://golang.org/dl/) (versi 1.21 atau lebih baru)
  * [Docker](https://www.docker.com/get-started) dan [Docker Compose](https://docs.docker.com/compose/install/)
  * [Make](https://www.gnu.org/software/make/)
  * [psql](https://www.postgresql.org/docs/current/app-psql.html) (opsional, untuk koneksi manual ke database)

### Instalasi

1.  **Clone repositori ini:**

    ```bash
    git clone https://github.com/username/gochi-boilerplate.git
    cd gochi-boilerplate
    ```

2.  **Konfigurasi Environment:**
    Salin file `.env.example` menjadi `.env` dan sesuaikan nilainya jika perlu.

    ```bash
    cp .env.example .env
    ```

3.  **Instal Dependensi Go:**
    Perintah ini akan mengunduh semua library yang dibutuhkan.

    ```bash
    go mod tidy
    ```

4.  **Jalankan Database:**
    Perintah ini akan memulai container PostgreSQL di latar belakang menggunakan Docker Compose.

    ```bash
    make db-up
    ```

5.  **Jalankan Migrasi Database:**
    Perintah ini akan membuat tabel `users` dan `products` sesuai skema di file `.sql`.

    ```bash
    make db-migrate
    ```

-----

## 📦 Penggunaan

### Menjalankan Server Development

Untuk menjalankan server dengan *hot-reload* (memerlukan instalasi `air`), atau cukup jalankan dengan perintah standar:

```bash
make run
```

Server akan berjalan di `http://localhost:8080`.

### Membangun Binary untuk Produksi

Untuk meng-kompilasi aplikasi menjadi satu file *binary* yang siap di-deploy:

```bash
make build
```

Hasilnya akan berada di folder `bin/`.

-----

## 📜 Daftar Perintah `Makefile`

| Perintah         | Deskripsi                                                                |
| ---------------- | ------------------------------------------------------------------------ |
| `make run`         | Menjalankan aplikasi Go dalam mode development.                          |
| `make build`       | Meng-kompilasi aplikasi menjadi file binary di folder `bin/`.            |
| `make test`        | Menjalankan semua unit test di dalam proyek.                             |
| `make clean`       | Menghapus artefak hasil build dari folder `bin/`.                         |
| `make tidy`        | Merapikan dependensi di `go.mod`.                                        |
| `make swag`        | Men-generate atau memperbarui dokumentasi Swagger di folder `docs/`.      |
| `make db-up`       | Menjalankan container database PostgreSQL dengan Docker Compose.         |
| `make db-down`     | Menghentikan dan menghapus container database.                           |
| `make db-migrate`  | Menjalankan skrip migrasi SQL ke database.                               |
| `make db-connect`  | Membuka shell `psql` interaktif ke dalam container database.             |
| `make help`        | Menampilkan daftar semua perintah yang tersedia.                         |

-----

## 📖 Endpoint API

Dokumentasi API yang lengkap dan interaktif tersedia melalui **Swagger UI**. Setelah server berjalan, buka URL berikut di browser Anda:

➡️ **[http://localhost:8080/swagger/index.html](https://www.google.com/search?q=http://localhost:8080/swagger/index.html)**

### Ringkasan Endpoint

#### Autentikasi

| Metode | Path             | Deskripsi                          |
| ------ | ---------------- | ---------------------------------- |
| `POST` | `/auth/register` | Mendaftarkan pengguna baru.        |
| `POST` | `/auth/login`    | Login untuk mendapatkan token JWT. |

#### Produk (Memerlukan Autentikasi)

| Metode   | Path             | Deskripsi                                |
| -------- | ---------------- | ---------------------------------------- |
| `POST`   | `/products`      | Membuat produk baru.                     |
| `GET`    | `/products`      | Mendapatkan daftar semua produk.         |
| `GET`    | `/products/{id}` | Mendapatkan detail satu produk.          |
| `PUT`    | `/products/{id}` | Memperbarui produk (memerlukan hak akses). |
| `DELETE` | `/products/{id}` | Menghapus produk (memerlukan hak akses).   |