# 🚀 Go User API

A production-ready REST API built with **Go (Fiber)**, **PostgreSQL**, and **SQLC**.  
This project demonstrates a clean architecture backend with real-world patterns like service layering, type-safe queries, validation, logging, and testing.

---

## ⚡ Tech Stack

🟢 Go (Golang) — backend language  
⚡ Fiber — fast HTTP framework  
🐘 PostgreSQL — relational database  
🧩 SQLC — type-safe SQL generator  
🪵 Zap — structured logging  
✅ Validator.v10 — request validation  

---

## ✨ Features

🔥 Full CRUD operations (Create, Read, Update, Delete)  
📄 Pagination support (limit & offset)  
🧠 Clean Architecture (Handler → Service → Repository)  
🎂 Age calculation from Date of Birth  
🛡️ Input validation for safe requests  
📊 Structured logs using Zap  
🧪 Unit tested handler & service layers  
⚡ Type-safe database queries using SQLC  

---

## 🏗️ Project Architecture
# 🚀 Go User API

A production-ready REST API built with **Go (Fiber)**, **PostgreSQL**, and **SQLC**.  
This project demonstrates a clean architecture backend with real-world patterns like service layering, type-safe queries, validation, logging, and testing.

---

## ⚡ Tech Stack

🟢 Go (Golang) — backend language  
⚡ Fiber — fast HTTP framework  
🐘 PostgreSQL — relational database  
🧩 SQLC — type-safe SQL generator  
🪵 Zap — structured logging  
✅ Validator.v10 — request validation  

---

## ✨ Features

🔥 Full CRUD operations (Create, Read, Update, Delete)  
📄 Pagination support (limit & offset)  
🧠 Clean Architecture (Handler → Service → Repository)  
🎂 Age calculation from Date of Birth  
🛡️ Input validation for safe requests  
📊 Structured logs using Zap  
🧪 Unit tested handler & service layers  
⚡ Type-safe database queries using SQLC  

---

## 🏗️ Project Architecture
# 🚀 Go User API

A production-ready REST API built with **Go (Fiber)**, **PostgreSQL**, and **SQLC**.  
This project demonstrates a clean architecture backend with real-world patterns like service layering, type-safe queries, validation, logging, and testing.

---

## ⚡ Tech Stack

🟢 Go (Golang) — backend language  
⚡ Fiber — fast HTTP framework  
🐘 PostgreSQL — relational database  
🧩 SQLC — type-safe SQL generator  
🪵 Zap — structured logging  
✅ Validator.v10 — request validation  

---

## ✨ Features

🔥 Full CRUD operations (Create, Read, Update, Delete)  
📄 Pagination support (limit & offset)  
🧠 Clean Architecture (Handler → Service → Repository)  
🎂 Age calculation from Date of Birth  
🛡️ Input validation for safe requests  
📊 Structured logs using Zap  
🧪 Unit tested handler & service layers  
⚡ Type-safe database queries using SQLC  

---

## 🏗️ Project Architecture

cmd/server → App entry point 🚀
db/migrations → Database schema 📦
db/sqlc → Generated SQL queries 🧩
internal/handler → HTTP layer 🌐
internal/service → Business logic 🧠
internal/repository → Database layer 🗄️
internal/logger → Logging system 🪵
internal/routes → API routes 🛣️


---

## ⚙️ Setup Guide

### 1️⃣ Clone the repo

git clone https://github.com/your-username/go-user-api.git
cd go-user-api

2️⃣ Create database

CREATE DATABASE userdb;

3️⃣ Run migrations

migrate -path db/migrations -database "postgres://postgres:password@localhost:5432/userdb?sslmode=disable" up

4️⃣ Start server 🚀

go run cmd/server/main.go

Server runs at:

http://localhost:3000

📡 API Endpoints

➕ Create User

POST /users
```bash
{
  "name": "John",
  "dob": "1990-01-01"
}
---
🔍 Get User
```bash
GET /users/:id
---

📃 List Users
```bash
GET /users?page=1&limit=10
---
✏️ Update User
```bash
PUT /users/:id
{
  "name": "John Updated",
  "dob": "1991-01-01"
}
---
❌ Delete User
```bash
DELETE /users/:id

🧪 Testing

Run all tests:
```bash 
go test ./...

🧠 What makes this project special?

💡 Built with real-world backend architecture
💡 Clean separation of concerns
💡 Fully testable structure
💡 Production-style logging
💡 Scalable folder design
💡 SQLC ensures zero raw SQL runtime errors

👨‍💻 Author

Built as a backend engineering project using Go, PostgreSQL, and SQLC.
