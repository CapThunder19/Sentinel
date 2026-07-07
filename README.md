# Sentinel

> Cloud-Native Distributed Backend Platform built with **Go** and **Rust**.

Sentinel is a production-inspired distributed backend system designed to explore microservices, authentication, distributed systems, cloud-native technologies, and high-performance networking.

---

# Tech Stack

### Backend
- Go
- Gin
- PostgreSQL
- pgx
- bcrypt
- JWT
- Docker

### Planned
- Rust (API Gateway)
- Redis
- Kafka
- gRPC
- Kubernetes
- Prometheus
- Grafana

---

# Current Architecture

```
                Client
                   │
                   ▼
            Auth Service (Go)
                   │
         ┌─────────┴─────────┐
         ▼                   ▼
    Auth Handler        JWT Service
         │
         ▼
     Auth Service
         │
         ▼
 User Repository
         │
         ▼
   PostgreSQL
```

---

# Features

## Authentication

- User Registration
- User Login
- Password Hashing (bcrypt)
- JWT Access Token Generation
- Repository Pattern
- Dependency Injection
- Service Layer
- REST APIs

---

# API Endpoints

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

Response

```json
{
    "id": "...",
    "username": "anirudh",
    "email": "ani@example.com"
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
    "access_token": "<jwt>",
    "token_type": "Bearer"
}
```

---

# Project Structure

```
backend/
│
├── auth-service/
│   ├── cmd/
│   ├── internal/
│   │   ├── handler/
│   │   ├── repository/
│   │   ├── routes/
│   │   ├── service/
│   │   └── models/
│
├── shared/
│   ├── auth/
│   ├── config/
│   ├── database/
│   └── logger/
│
└── migrations/
```

---

# Testing

### Integration Tests

- PostgreSQL Repository Tests

### Unit Tests

- Registration Service
- Login Service
- JWT Generation

Run all tests

```bash
go test ./...
```

---

# Database

PostgreSQL is managed through Docker.

Run migrations

```bash
migrate \
-path migrations \
-database "postgres://postgres:<password>@localhost:5433/sentinel?sslmode=disable" \
up
```

---

# Current Progress

- [x] Docker Setup
- [x] PostgreSQL Integration
- [x] Database Migrations
- [x] Repository Layer
- [x] Repository Tests
- [x] Service Layer
- [x] Service Tests
- [x] User Registration
- [x] User Login
- [x] Password Hashing
- [x] JWT Authentication
- [ ] JWT Middleware
- [ ] Refresh Tokens
- [ ] Rust API Gateway
- [ ] Redis Caching
- [ ] Kafka Messaging
- [ ] gRPC Communication
- [ ] Kubernetes Deployment
- [ ] Prometheus Metrics
- [ ] Grafana Dashboards

---

# Roadmap

### Phase 1 ✅

- Authentication Service
- PostgreSQL
- JWT
- Testing

### Phase 2

- Rust API Gateway
- Authentication Middleware
- Protected APIs

### Phase 3

- User Service
- Redis
- Kafka
- gRPC

### Phase 4

- Kubernetes
- Monitoring
- CI/CD