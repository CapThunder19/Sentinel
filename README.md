# Sentinel

> A cloud-native distributed backend platform built with **Rust** and **Go**, designed to explore modern backend engineering, microservices, distributed systems, and cloud-native technologies.

---

# Architecture

```text
                    Client / Frontend
                           │
                           ▼
                  🦀 Rust API Gateway
        ┌──────────────────────────────────┐
        │ Generic Reverse Proxy            │
        │ Dynamic Routing                  │
        │ Header Forwarding                │
        │ Request Forwarding               │
        └──────────────────────────────────┘
                           │
                           ▼
                  🐹 Go Auth Service
                           │
                    PostgreSQL Database
```

---

# Tech Stack

## Gateway

- Rust
- Axum
- Tokio
- Reqwest

## Auth Service

- Go
- Gin
- PostgreSQL
- pgx
- bcrypt
- JWT

## Database

- PostgreSQL

---

# Features

## Rust API Gateway

- Generic Reverse Proxy
- Dynamic Route Forwarding (`/auth/*`)
- Request Body Forwarding
- Header Forwarding
- Health Endpoint
- Shared HTTP Client
- Configuration Management
- Service-to-Service Communication

---

## Go Authentication Service

- User Registration
- User Login
- Password Hashing (bcrypt)
- JWT Authentication
- Protected `/me` Endpoint
- PostgreSQL Integration
- Repository Pattern
- Service Layer
- Unit Tests
- Repository Tests

---

# Project Structure

```
sentinel/

├── gateway/                 # Rust API Gateway
│   ├── src/
│   │   ├── handlers/
│   │   ├── routes.rs
│   │   ├── config.rs
│   │   ├── state.rs
│   │   └── main.rs
│   └── Cargo.toml
│
├── backend/
│   ├── auth-service/
│   │   ├── cmd/
│   │   ├── internal/
│   │   │   ├── handler/
│   │   │   ├── repository/
│   │   │   ├── routes/
│   │   │   ├── service/
│   │   │   └── models/
│   │   └── migrations/
│   │
│   └── shared/
│       ├── auth/
│       ├── config/
│       ├── database/
│       └── logger/
│
└── README.md
```

---

# Authentication Flow

```text
Client
   │
POST /auth/login
   │
   ▼
Rust Gateway
   │
Generic Reverse Proxy
   │
POST /login
   │
   ▼
Go Auth Service
   │
Verify Credentials
   │
Generate JWT
   │
   ▼
Rust Gateway
   │
   ▼
Client
```

---

# Protected Route Flow

```text
Client
   │
Authorization: Bearer <JWT>
   │
   ▼
Rust Gateway
   │
Header Forwarding
   │
   ▼
Go Auth Service
   │
JWT Middleware
   │
Repository
   │
PostgreSQL
   │
   ▼
User Profile
```

---

# Current Endpoints

## Gateway

| Method | Endpoint |
|---------|----------|
| GET | `/health` |
| ANY | `/auth/*` |

---

## Auth Service

| Method | Endpoint | Description |
|---------|----------|-------------|
| GET | `/` | Health Check |
| POST | `/register` | Register User |
| POST | `/login` | Login |
| GET | `/me` | Current User |

---

# Current Progress

## Completed

### Infrastructure

- [x] Project Structure
- [x] Shared Configuration
- [x] Logging
- [x] PostgreSQL Integration

### Rust Gateway

- [x] Axum Setup
- [x] Generic Reverse Proxy
- [x] Dynamic Route Matching
- [x] Header Forwarding
- [x] Body Forwarding
- [x] Health Endpoint
- [x] Service-to-Service Communication

### Go Auth Service

- [x] Registration
- [x] Login
- [x] JWT Authentication
- [x] Protected Routes
- [x] Repository Pattern
- [x] Unit Tests
- [x] Repository Tests

---

# Upcoming Roadmap

## User Service

- [ ] User Profiles
- [ ] Update Profile
- [ ] Search Users

---

## Gateway

- [ ] Query Parameter Forwarding
- [ ] Response Header Forwarding
- [ ] Request Logging
- [ ] CORS
- [ ] Rate Limiting
- [ ] Metrics
- [ ] Tracing

---

## Platform

- [ ] Redis
- [ ] gRPC
- [ ] Kafka
- [ ] Docker
- [ ] Kubernetes
- [ ] Prometheus
- [ ] Grafana

---

## Frontend

- [ ] Next.js
- [ ] TypeScript
- [ ] Tailwind CSS
- [ ] Authentication
- [ ] Dashboard

---

# Learning Goals

Sentinel is built to gain hands-on experience with:

- Rust Backend Development
- Go Microservices
- Distributed Systems
- API Gateway Design
- Authentication & Authorization
- Cloud-Native Development
- System Design
- Docker & Kubernetes
- Event-Driven Architecture
- Observability
- Production Backend Engineering

---

# Status

**Current Phase:** Gateway & Authentication ✅

The project currently consists of a production-style Rust API Gateway routing requests to a Go-based Authentication Service backed by PostgreSQL. Authentication flows, protected routes, and generic reverse proxying are fully functional.

The next milestone is expanding the platform with additional microservices, beginning with a dedicated User Service.