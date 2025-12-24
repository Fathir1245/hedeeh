# 🚀 hedeeh – Go Project Scaffolding CLI

**hedeeh** adalah **CLI tool berbasis Go** untuk menghasilkan (*scaffold*) struktur project backend Go secara otomatis dengan pilihan:

* 📦 Database: **MySQL**
* 🌐 Router: **Chi / Gin / Standard net/http**
* 🧩 Arsitektur: Clean Architecture (handler → service → repository)
* 🧠 Template **embedded ke binary** (tidak bergantung file eksternal)

Cocok untuk:

* Developer Go
* Mahasiswa
* Backend engineer
* Boilerplate cepat untuk API service

---

## ✨ Fitur Utama

* ✅ Generate project Go siap pakai
* ✅ Pilihan router (Chi / Gin / net/http)
* ✅ Konfigurasi database MySQL
* ✅ Struktur folder rapi & konsisten
* ✅ Auto `go mod init` & `go mod tidy`
* ✅ Template di-*embed* ke binary (portable)
* ✅ Bisa dijalankan di mesin mana pun

---

## 📁 Struktur Project yang Dihasilkan

```bash
project-name/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── database/
│   │   └── db.go
│   ├── model/
│   │   └── models.go
│   ├── repository/
│   │   └── repository.go
│   ├── service/
│   │   └── service.go
│   ├── handler/
│   │   └── handler.go
│   └── router/
│       └── router.go
├── .env
├── .gitignore
├── Dockerfile
├── docker-compose.yml
└── go.mod
```

---

## 🛠️ Instalasi

### Clone Repository

```bash
git clone https://github.com/username/hedeeh.git
cd hedeeh
```

### Build Binary

```bash
go build -o hedeeh
```

---

## ▶️ Cara Menggunakan

Jalankan perintah berikut:

```bash
./hedeeh init
```

Lalu ikuti prompt interaktif:

```text
? Apa nama projek kamu? my-api
? Pilih Bahasa Pemrograman: Go
? Pilih Database Utama: MySQL
? Pilih Router/Framework: Chi
```

Hasilnya:

```text
🚀 Oke, 'hedeeh' bakal buatin projek my-api...
🛠  Sedang membangun project: my-api ...
✅ Created: cmd/api/main.go
✅ Created: internal/router/router.go
📦 Menginisialisasi Go Module...
```

---

## 🧠 Cara Kerja Template

* Semua file `.tpl` **di-embed ke binary** menggunakan `embed.FS`
* Tidak ada dependensi ke file eksternal
* Path template diakses langsung dari memory

```go
//go:embed go/**
var FS embed.FS
```

Ini membuat `hedeeh`:

* portable
* aman dibagikan
* siap dipakai di mana saja

---

## ⚙️ Dukungan Router

| Router              | Status |
| ------------------- | ------ |
| net/http (Standard) | ✅      |
| Chi                 | ✅      |
| Gin                 | ✅      |

---

## 📦 Database

| Database   | Status       |
| ---------- | ------------ |
| MySQL      | ✅            |
| PostgreSQL | 🔜 (planned) |

---

## 🧪 Requirement

* Go ≥ 1.20
* OS: Linux / macOS / Windows

---

## 🧩 Roadmap

* [ ] PostgreSQL support
* [ ] CRUD generator (`hedeeh add crud user`)
* [ ] Middleware generator
* [ ] Auth template (JWT)
* [ ] AI-assisted code injection

---

## 📄 Lisensi

MIT License
Bebas digunakan untuk keperluan pribadi maupun komersial.

---

## 👤 Author

**Moh. Faathir Ash Shaff**
Computer Science Student & Backend Developer

---
