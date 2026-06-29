# 🗺️ My Journey to Become a Professional Backend Developer

Repositori ini berfungsi sebagai jurnal, pelacak progres (_progress tracker_), dan tempat penyimpanan kode latihan saya selama mengikuti roadmap pembelajaran Backend Engineering.

## 🛠️ My Local Development Environment

- **OS:** Arch Linux 🐧
- **Editor:** Neovim (LazyVim v14+) 🚀
- **Core Stacks:** Node.js/TypeScript, Golang, .NET (C#)
- **Database:** MariaDB

---

## 🛣️ The Detailed Backend Roadmap

### 📦 Phase 1: Programming Language Fundamentals & Logic

Pilih salah satu bahasa terlebih dahulu sebagai fondasi utama sebelum pindah ke bahasa berikutnya.

#### 🟦 1. JavaScript & TypeScript (Node.js Run-time)

- [ ] **Basic Syntax:** Variabel (`let`, `const`), Tipe Data Primitif, Operator, Control Flow (`if-else`, `switch-case`, `for`, `while`).
- [ ] **Data Structures:** Manipulasi Array (Method `.map()`, `.filter()`, `.reduce()`), Object Literals, dan JSON Parsing.
- [ ] **Asynchronous Programming:** Memahami _Event Loop_, Callback Hell, `Promises` (`.then()`, `.catch()`), dan sintaks modern `async/await`.
- [ ] **TypeScript Type-Safety:** Mengubah proyek JS menjadi TS, memahami _Static Typing_, penggunaan `Interface` vs `Type`, `Enums`, dan _Generics_.

#### 🟫 2. Go (Golang)

- [ ] **The Zen of Go:** Karakteristik kompilasi Go, instalasi SDK di Arch Linux, pembuatan modul (`go mod init`).
- [ ] **Syntax & Types:** Variabel, tipe data, _zero values_, _implicit conversion_, kontrol alur tanpa tanda kurung.
- [ ] **Pointers & Memory:** Memahami alokasi memori alamat (`&`) dan dereferensi (`*`) — sangat krusial untuk efisiensi performa backend Go.
- [ ] **Composite Types:** Perbedaan Array statis dan _Slices_ dinamis, penggunaan _Maps_ (key-value storage).
- [ ] **OOP ala Go:** Membuat struktur data dengan `Structs`, menambahkan fungsi melalui _Method Receivers_, dan abstraksi dengan `Interfaces`.
- [ ] **Concurrency Basics:** Pengenalan `Goroutines` untuk multi-threading ringan dan `Channels` untuk komunikasi data antar thread.

#### 🟪 3. C# (.NET Ecosystem)

- [ ] **The .NET CLI:** Membuat proyek via terminal (`dotnet new console`, `dotnet run`), manajemen file `.csproj`.
- [ ] **Strongly-Typed Foundations:** Tipe data bernilai (_Value Types_) vs Tipe data referensi (_Reference Types_), _nullability_.
- [ ] **Deep OOP (Object-Oriented Programming):**
  - Encapsulation (Access Modifiers: `public`, `private`, `protected`, `internal`).
  - Inheritance (Pewarisan _class_, penggunaan keyword `base`).
  - Polymorphism (`virtual`, `override`, `abstract` classes).
- [ ] **Advanced C# Expressions:** Koleksi data generik (`List<T>`, `Dictionary<TKey, TValue>`), delegasi fungsi, lambda expressions, dan sintaks **LINQ** (Language Integrated Query).

---

### 🗄️ Phase 2: Relational Database Mastery (MariaDB)

Fokus pada cara merancang, menyimpan, mengamankan, dan mengambil data secara optimal.

- [ ] **Database Setup via CLI:** Instalasi MariaDB di Arch, setup user, manajemen hak akses (`GRANT ALL PRIVILEGES`), dan koneksi via terminal terminal (`mariadb -u root -p`).
- [ ] **DDL (Data Definition Language):** Membuat arsitektur database (`CREATE DATABASE`), tabel, relasi, batasan data (`NOT NULL`, `UNIQUE`, `CHECK`).
- [ ] **DML (Data Manipulation Language):** Operasi CRUD dasar (`INSERT`, `SELECT` dengan filter `WHERE`, `UPDATE`, `DELETE`).
- [ ] **Relational Engineering:**
  - Menghubungkan tabel dengan _Primary Key_ dan _Foreign Key_.
  - Menguasai teknik penggabungan data: `INNER JOIN`, `LEFT JOIN`, `RIGHT JOIN`.
  - Normalisasi data (1NF, 2NF, 3NF) untuk menghindari data redundan/duplikat.
- [ ] **Advanced Queries:** Agregasi data (`COUNT`, `SUM`, `AVG`, `GROUP BY`, `HAVING`), sub-queries (query di dalam query).
- [ ] **Performance & Integrity:**
  - Pembuatan _Indexes_ untuk optimasi kecepatan `SELECT` pada baris data skala besar.
  - Penerapan Konsep **ACID** melalui database _Transactions_ (`START TRANSACTION`, `COMMIT`, `ROLLBACK`) untuk mencegah kegagalan mutasi data di tengah jalan.

---

### 🌐 Phase 3: RESTful Web API Architecture & Frameworks

Menghubungkan kode backend dengan database MariaDB agar bisa dikonsumsi oleh aplikasi luar lewat protokol internet (HTTP).

#### 🌐 1. HTTP Protocol Concepts

- [ ] Memahami siklus Request-Response, HTTP Methods (`GET`, `POST`, `PUT`, `DELETE`, `PATCH`).
- [ ] Memahami HTTP Status Codes (`200 OK`, `201 Created`, `400 Bad Request`, `401 Unauthorized`, `404 Not Found`, `500 Internal Server Error`).

#### 🛠️ 2. Hands-on Frameworks & Integration

- [ ] **Node.js:** Membuat REST API dengan **Express.js** atau **NestJS**, menggunakan **Prisma ORM** untuk terhubung ke MariaDB.
- [ ] **Go:** Membuat REST API berkecepatan tinggi dengan **Fiber** atau **Gin**, menggunakan **GORM** atau _Raw SQL queries_ bawaan paket `database/sql`.
- [ ] **C#:** Membuat arsitektur Web API terstruktur dengan **ASP.NET Core**, memanfaatkan **Entity Framework Core (EF Core)** untuk pemetaan database MariaDB secara otomatis lewat kode C#.

---

### 🔒 Phase 4: Security, Authentication, & Testing

Tahap krusial untuk menjamin aplikasi backend aman dari serangan luar dan bebas dari bug logika.

- [ ] **Password Security:** Menerapkan teknik enkripsi satu arah (_Password Hashing_) menggunakan algoritma aman seperti `bcrypt` atau `argon2`.
- [ ] **Authentication Mechanism:** Implementasi Stateless Authentication menggunakan **JWT (JSON Web Tokens)**: pembuatan token saat login, dan validasi token lewat _Middleware_ di setiap endpoint terproteksi.
- [ ] **Authorization (RBAC):** Membatasi akses user berdasarkan peran (_Role-Based Access Control_, contoh: memisahkan hak akses Route antara `User` biasa dan `Admin`).
- [ ] **Web Security Standard:** Mengatasi isu keamanan seperti pencegahan _SQL Injection_ (menggunakan _Parameterized Queries_ / ORM) dan konfigurasi _CORS_ (Cross-Origin Resource Sharing).
- [ ] **Automated Testing:** Menulis skenario pengujian unit (**Unit Testing**) menggunakan framework pengujian seperti `Jest` (Node.js), paket bawaan `testing` (Go), atau `xUnit` (C#).

---

### 🚀 Phase 5: Scalability, Infrastructure, & DevOps (The Handal Level)

Mempersiapkan aplikasi agar mampu menangani jutaan pengguna, terdistribusi dengan baik, dan siap di-deploy ke server produksi.

- [ ] **Caching Layer (Redis):** Mengintegrasikan Redis sebagai _In-Memory Database_ untuk menyimpan data respon yang jarang berubah, mengurangi beban query langsung ke MariaDB.
- [ ] **Containerization (Docker):** Membuat file `Dockerfile` untuk membungkus kode backend dan file `docker-compose.yml` untuk menjalankan aplikasi beserta kontainer MariaDB secara instan di mana saja.
- [ ] **Message Brokers:** Mempelajari arsitektur _Event-Driven_ menggunakan **RabbitMQ** untuk memproses antrean tugas berat di latar belakang secara asinkronus (_Asynchronous background worker_).
- [ ] **CI/CD Pipelines:** Membuat skrip otomatisasi pengujian kode setiap kali melakukan `git push` menggunakan layanan **GitHub Actions**.
- [ ] **Cloud Deployment:** Cara melakukan sewa server Cloud VPS (seperti DigitalOcean, AWS, Linode), melakukan manajemen server via SSH Linux, dan mengonfigurasi **Nginx** sebagai _Reverse Proxy_ untuk mengarahkan domain publik ke port internal aplikasi backend-mu.

---

## 📈 Current Progress Tracker

> Perbarui persentase dan tanda centang di atas secara berkala setiap kali kamu menyelesaikan materi latihan!

- **Phase 1 (Language Fundamentals):** 0%
- **Phase 2 (Database Mastery):** 0%
- **Phase 3 (Web API Development):** 0%
- **Phase 4 (Security & Testing):** 0%
- **Phase 5 (Advanced & DevOps):** 0%

---

_“An investment in knowledge pays the best interest. Keep pushing, Arch-User!”_
