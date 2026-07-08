# Sentinel 🛡️

> A cloud-native distributed backend platform built with **Rust**, **Go**, **PostgreSQL**, **Redis**, **Kafka**, **gRPC**, **Docker**, **Kubernetes**, and **Next.js**.

Sentinel is a production-inspired microservices platform designed to learn and demonstrate modern backend architecture, distributed systems, and cloud-native development.

---

# 🚀 Tech Stack

## Backend
- Go (Gin)
- Rust (Axum + Tokio) *(Coming Soon)*
- PostgreSQL
- Redis *(Planned)*
- Kafka *(Planned)*
- gRPC *(Planned)*

## Frontend
- Next.js
- TypeScript
- Tailwind CSS

## DevOps
- Docker
- Kubernetes *(Planned)*
- Prometheus *(Planned)*
- Grafana *(Planned)*

---

# 📂 Project Structure

```
Sentinel/
│
├── backend/
│   ├── auth-service/
│   ├── shared/
│   └── migrations/
│
├── gateway/          # Rust API Gateway (Coming Soon)
│
├── frontend/         # Next.js Frontend (Coming Soon)
│
├── deployments/
│
└── docs/
```

---

# ✅ Current Features

## Authentication Service

- User Registration
- Secure Password Hashing (bcrypt)
- User Login
- JWT Authentication
- JWT Middleware
- Protected Routes
- Authenticated User Profile (`/me`)
- Repository Pattern
- Service Layer
- Dependency Injection
- PostgreSQL Integration
- Database Migrations
- Docker Support

---

# 🏗️ Authentication Architecture

```
               Client
                  │
                  ▼
         POST /register
                  │
             Hash Password
             (bcrypt)
                  │
                  ▼
             PostgreSQL
                  ▲
                  │
         POST /login
                  │
          Validate Password
                  │
                  ▼
            Generate JWT
                  │
                  ▼
         Return Access Token
                  │
                  ▼
Authorization: Bearer <JWT>
                  │
                  ▼
          JWT Middleware
                  │
        Extract User ID
                  │
                  ▼
          Repository Layer
                  │
                  ▼
            PostgreSQL
                  │
                  ▼
       Return User Profile
```

---

# 📌 API Endpoints

## Health Check

```
GET /
```

---

## Register

```
POST /register
```

Request

```json
{
  "username": "anirudh",
  "email": "ani@example.com",
  "password": "password123"
}
```

---

## Login

```
POST /login
```

Request

```json
{
  "email": "ani@example.com",
  "password": "password123"
}
```

Response

```json
{
  "AccessToken": "<jwt-token>",
  "TokenType": "Bearer"
}
```

---

## Authenticated User

```
GET /me
```

Headers

```
Authorization: Bearer <JWT>
```

Response

```json
{
  "id": "...",
  "username": "anirudh",
  "email": "ani@example.com",
  "role": "USER",
  "is_verified": false,
  "created_at": "...",
  "updated_at": "..."
}
```

---

# 🗄️ Database

Current tables

- users
- schema_migrations

---

# 🧪 Running Tests

```
go test ./...
```

Current Status

- Repository Tests ✅
- Service Tests ✅
- Authentication Tests ✅
- JWT Tests ✅
- Integration Tests ✅

---

# 🐳 Running Locally

Start PostgreSQL

```bash
cd deployments/docker

docker compose up -d
```

Run migrations

```bash
cd backend

migrate \
-path migrations \
-database "postgres://postgres:postgres@localhost:5433/sentinel?sslmode=disable" \
up
```

Run Auth Service

```bash
go run ./...
```

---

# 🛣️ Roadmap

## ✅ Phase 1

- [x] Project Setup
- [x] PostgreSQL
- [x] Docker
- [x] Migrations
- [x] Repository Layer
- [x] Service Layer
- [x] Registration
- [x] Login
- [x] JWT Authentication
- [x] JWT Middleware
- [x] Protected Routes
- [x] User Profile Endpoint

---

## 🚧 Phase 2

- [ ] Rust API Gateway
- [ ] Reverse Proxy
- [ ] Request Routing
- [ ] JWT Validation
- [ ] Rate Limiting
- [ ] Request Logging

---

## 🚧 Phase 3

- [ ] Next.js Frontend
- [ ] Dashboard
- [ ] Authentication UI
- [ ] Protected Pages

---

## 🚧 Phase 4

- [ ] User Service
- [ ] Redis
- [ ] Kafka
- [ ] gRPC

---

## 🚧 Phase 5

- [ ] Kubernetes
- [ ] Prometheus
- [ ] Grafana
- [ ] CI/CD

---

# 📈 Current Progress

```
Infrastructure        ████████████████████ 100%

Auth Service          ████████████████████ 100%

Rust Gateway          ░░░░░░░░░░░░░░░░░░░░   0%

Frontend              ░░░░░░░░░░░░░░░░░░░░   0%

User Service          ░░░░░░░░░░░░░░░░░░░░   0%

Redis                 ░░░░░░░░░░░░░░░░░░░░   0%

Kafka                 ░░░░░░░░░░░░░░░░░░░░   0%

gRPC                  ░░░░░░░░░░░░░░░░░░░░   0%

Kubernetes            ░░░░░░░░░░░░░░░░░░░░   0%
```

---

# 🎯 Project Goal

The goal of Sentinel is to build a production-inspired distributed backend platform that demonstrates:

- Clean Architecture
- Microservices
- Distributed Systems
- High-performance Rust Networking
- Cloud-native Development
- Modern DevOps Practices

---

## 👨‍💻 Author

**Anirudh Patwal**

