# Project Directory Guide

Dokumentasi fungsi setiap folder dalam project Go ERP Backend.

## Folder

| Folder | Fungsi |
|---|---|
| `cmd/` | Entry point aplikasi. Tempat file `main.go` untuk menjalankan aplikasi. |
| `internal/` | Menyimpan source code utama aplikasi yang bersifat internal. |
| `internal/platform/` | Menyimpan infrastruktur teknis seperti database, cache, logger, dan HTTP server. |
| `internal/modules/` | Menyimpan modul dan business logic ERP seperti authentication, inventory, sales, dan accounting. |
| `internal/shared/` | Menyimpan komponen umum yang digunakan bersama oleh berbagai modul. |
| `migrations/` | Menyimpan file SQL untuk perubahan struktur database. |
| `configs/` | Menyimpan konfigurasi aplikasi yang bukan informasi rahasia. |
| `docs/` | Menyimpan dokumentasi teknis, arsitektur, API, dan database. |
| `scripts/` | Menyimpan script untuk otomatisasi development dan deployment. |
| `tests/` | Menyimpan integration test dan end-to-end test. |

## Root Files

| File | Fungsi |
|---|---|
| `.env.example` | Contoh environment variables yang dibutuhkan aplikasi. |
| `.gitignore` | Menentukan file dan folder yang tidak dimasukkan ke Git. |
| `Dockerfile` | Instruksi untuk membangun Docker image aplikasi. |
| `docker-compose.yml` | Konfigurasi untuk menjalankan aplikasi dan service pendukung menggunakan Docker Compose. |
| `Makefile` | Menyediakan shortcut untuk perintah development. |
| `go.mod` | Mendefinisikan Go module dan dependency project. |
| `go.sum` | Menyimpan checksum dependency Go. |
| `LICENSE` | Menentukan ketentuan penggunaan dan distribusi software. |
| `README.md` | Dokumentasi utama project.