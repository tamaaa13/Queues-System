# 📋 Sistem Antrian Digital

Aplikasi sistem antrian digital berbasis web yang terdiri dari dua komponen utama:

- 📦 `backend/` – API server dengan Go (native SQL, Viper, Docker)
- 💻 `frontend/` – Antarmuka pengguna Vue.js untuk admin, super admin, dan display publik

---

## 🧠 Fitur Utama

### ✅ 1. **Pencetakan Tiket Antrian**

- Pengguna dapat memilih kategori layanan
- Sistem mencetak tiket antrian dengan nomor unik dan waktu.

### ✅ 2. **Pemanggilan Antrian**

- Admin Loket dapat memanggil nomor antrian berdasarkan kategori yang ditangani.
- Tersedia tampilan layar publik yang menunjukkan nomor yang sedang dipanggil dan ke loket mana.
- Pemanggilan dilengkapi dengan **suara otomatis (text-to-speech)**.

### ✅ 3. **Manajemen Admin (Super Admin)**

- Super Admin dapat:
  - Menambahkan admin loket.
  - Menghapus atau mengedit akun admin.
  - Melihat seluruh riwayat antrian yang dipanggil oleh setiap admin.

### ✅ 4. **Tampilan Realtime**

- Tampilan publik menampilkan:
  - Nomor antrian yang sedang dipanggil.
  - Loket tujuan.
  - Kategori layanan.
- Real-time update via WebSocket (opsional).

---

## ⚙️ Teknologi yang Digunakan

### Backend:

- 🧠 Go (Gin Framework)
- 📄 Native SQL (tanpa ORM)
- ⚙️ Viper (untuk load konfigurasi `.env` atau `yaml`)
- 🐬 MySQL
- 🔐 JWT Authentication
- 🔊 Text-to-Speech lokal (misalnya: eSpeak)
- 🐳 Docker (containerized deployment)

### Frontend:

- 🌐 Vue.js
- 🎨 UI responsif untuk touchscreen & TV
- 🌍 Native `fetch()` untuk komunikasi dengan API
- 🔊 Audio pemanggilan antrian

---

## 🚀 Cara Menjalankan

### 1. Clone Repository

```bash
git clone https://github.com/username/Queue.git
cd Queue
```

### 2. Siapkan File `.env`

Rename file .env.example menjadi .env dan sesuaikan dengan konfigurasi pribadi

### 3. Running Docker For Backend

- Pastikan sudah install docker dan docker compose
- Lalu jalankan docker

```bash
cd backend
docker-compose up --build
```

### 4. Running Frontend `(Vue JS)`

```bash
cd frontend
npm install
npm run dev
```
