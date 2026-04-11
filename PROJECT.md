# POS GO PROJECT (DDD)

## 🎯 Goal
Membangun POS backend dengan DDD, clean workflow, dan production-ready per fitur.

## 📦 Current Feature
- Product (Inventory)

## ✅ DONE
- [x] Setup project
- [x] Gin server
- [x] Product Entity (Domain)
- [x] Repository Interface
- [x] Use Case (Create & Get Product)
- [x] Repository Implementation (GORM)
- [x] Database connection & migration

## 🔄 IN PROGRESS
- [ ] Delivery Layer (HTTP Handler & Router)

## ⏭️ NEXT
- [ ] Test API (manual: Postman / curl)
- [ ] Tambahkan response DTO (jangan expose domain langsung)
- [ ] Error handling yang lebih rapi

## 🧠 NOTES
- Semua logic harus mulai dari domain
- Jangan loncat ke handler
- Domain tidak boleh kena JSON / HTTP
- Mapping domain ↔ DB sudah aman
- Usecase tetap bersih tanpa dependency luar
