# Sentinel

> A Cloud-Native Distributed Backend Platform built with Go, Rust, Docker, Kubernetes, PostgreSQL, Redis, Kafka, gRPC, and Next.js.

> **Status:** 🚧 In Development

---

## 📖 Overview

Sentinel is a production-grade distributed backend platform built to explore modern backend engineering and cloud-native technologies.

The goal is not to build another CRUD application, but to understand how large-scale backend systems are designed, deployed, monitored, and scaled.

This project focuses on:

- Clean Architecture
- Microservices
- Distributed Systems
- Event-Driven Architecture
- Cloud-Native Development
- High Performance Backend Engineering

---

# 🚀 Tech Stack

## Backend

- Go
- Gin
- pgx
- Zap Logger

## Database

- PostgreSQL

## Infrastructure

- Docker
- Docker Compose

## Frontend (Planned)

- Next.js
- TypeScript
- Tailwind CSS

## Upcoming Technologies

- Rust
- Redis
- Kafka
- gRPC
- Kubernetes
- Prometheus
- Grafana

---

# 📁 Project Structure

```text
Sentinel/

├── backend/
│   ├── auth-service/
│   └── shared/
│       ├── config/
│       ├── database/
│       └── logger/
│
├── frontend/
│
├── deployments/
│   ├── docker/
│   └── kubernetes/
│
├── docs/
│
└── README.md
```

---

# ✅ Current Progress

## Infrastructure

- [x] Repository setup
- [x] Project structure
- [x] Docker Compose
- [x] PostgreSQL Container
- [x] Persistent Docker Volume

---

## Backend

- [x] Go Module
- [x] Configuration Management
- [x] Zap Logging
- [x] PostgreSQL Connection Pool
- [x] Gin Server
- [x] Health Check Endpoint
- [x] Clean Project Structure

---

## Authentication

- [ ] Database Migrations
- [ ] User Model
- [ ] User Repository
- [ ] Password Hashing
- [ ] Register API
- [ ] Login API
- [ ] JWT Authentication
- [ ] Refresh Tokens
- [ ] Role-Based Access Control

---

## Services

- [ ] API Gateway (Rust)
- [ ] Auth Service
- [ ] User Service
- [ ] Project Service
- [ ] Notification Service

---

## Infrastructure (Upcoming)

- [ ] Redis
- [ ] Kafka
- [ ] gRPC
- [ ] Kubernetes
- [ ] Prometheus
- [ ] Grafana

---

# 🏗 Current Architecture

```
                 Docker
                    │
                    ▼
          PostgreSQL Container
                    │
                    ▼
           Connection Pool (pgx)
                    │
                    ▼
               Auth Service
                    │
                    ▼
                 Gin Router
                    │
                    ▼
                 HTTP Client
```

---

# ⚙️ Running Locally

## Clone

```bash
git clone https://github.com/CapThunder19/Sentinel.git

cd Sentinel
```

---

## Start PostgreSQL

```bash
cd deployments/docker

docker compose up -d
```

---

## Start Backend

```bash
cd backend

go run ./auth-service/cmd/server
```

---

## Verify

Visit:

```
http://localhost:8081/
```

Expected response:

```json
{
    "service":"auth-service",
    "status":"running"
}
```

---

# 🎯 Roadmap

Phase 1 ✅

- Project Setup
- Docker
- PostgreSQL
- Logging
- Configuration

Phase 2 🚧

- Database Migrations
- User Repository
- Authentication
- JWT

Phase 3

- Redis
- Caching
- Session Management

Phase 4

- Kafka
- Event Driven Architecture

Phase 5

- Rust API Gateway

Phase 6

- gRPC Communication

Phase 7

- Kubernetes Deployment

Phase 8

- Monitoring
- Prometheus
- Grafana

---

# 📚 Learning Goals

This project is designed to learn and practice:

- Go Backend Development
- Rust Systems Programming
- Distributed Systems
- Cloud Native Architecture
- Docker
- Kubernetes
- PostgreSQL
- Redis
- Kafka
- gRPC
- Authentication
- Observability

---

# 📜 License

MIT License