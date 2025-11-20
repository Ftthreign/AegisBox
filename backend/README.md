# AegisBox Backend

A secure and modular backend service for **AegisBox**, a modern password manager designed with a strong focus on **security**, **data integrity**, and **auditability**.  
Built with **Go**, **GORM**, **PostgreSQL**, and a fully structured **Clean Architecture** for maintainability and scalability.

---

## 🚀 Tech Stack

- **Go 1.22+**
- **GORM ORM**
- **PostgreSQL**
- **Argon2id password hashing**
- **Secure session store (Postgres-based)**
- **JSONB audit logging**
- **Rate limiting and login lockout**
- **Docker & Docker Compose**

---

## 📦 Docker Support

AegisBox backend includes a production-ready Dockerfile.

### Build image:

```sh
docker build -t aegisbox-backend .
```

### Run with Docker Compose:

```sh
docker compose up --build
```

Backend available at:

```
http://localhost:8080
```

---

## 📁 Project Structure

The backend follows **Clean Architecture**, dividing logic into domain, infrastructure, interface, and configuration layers.

```
backend/
├── Dockerfile
├── go.mod
├── go.sum
├── .env.example
│
├── cmd/
│   └── server/
│       └── main.go               # Application entry point
│
└── internal/
    ├── app/                      # Dependency injection container
    │   └── app.go
    │
    ├── config/                   # Modular configuration
    │   ├── config.go
    │   ├── database.go
    │   ├── security.go
    │   ├── session.go
    │   ├── limiter.go
    │   └── server.go
    │
    ├── domain/                   # Core business logic (entities & interfaces)
    │   ├── user/
    │   ├── credential/
    │   └── audit/
    │
    ├── infrastructure/           # Implementation layer
    │   ├── database/
    │   │   ├── postgres.go
    │   │   └── model/
    │   │       ├── user_model.go
    │   │       ├── credential_model.go
    │   │       └── audit_log_model.go
    │   └── repository/
    │       └── postgres/
    │
    └── interface/                # HTTP entry layer
        └── http/
            ├── handler/
            ├── dto/
            ├── middleware/
            └── router.go
```

---

## ⚙️ Environment Variables

Example `.env.example`:

```
# Database
DATABASE_URL=
DB_HOST=postgres
DB_PORT=5432
DB_USER=aegis
DB_PASSWORD=aegis
DB_NAME=aegisdb

# Server
PORT=8080
ENVIRONMENT=development
ALLOWED_ORIGINS=http://localhost:5173

# Session
SESSION_SECRET=your_very_secure_secret_key_at_least_32_chars
SESSION_MAX_AGE=1800
SESSION_IDLE_TIMEOUT=900

# Argon2 Hashing
ARGON2_TIME_COST=4
ARGON2_MEMORY_COST=65536
ARGON2_PARALLELISM=4
ARGON2_KEY_LENGTH=32

# Rate Limiting
LOGIN_RATE_LIMIT_MAX_ATTEMPTS=10
LOGIN_RATE_LIMIT_WINDOW=3600
LOGIN_LOCKOUT_DURATION=300

# Password Policy
MIN_PASSWORD_LENGTH=14
REQUIRE_UPPERCASE=true
REQUIRE_LOWERCASE=true
REQUIRE_NUMBERS=true
REQUIRE_SPECIAL_CHARS=true
```

> **Important:** `SESSION_SECRET` must be at least **32 characters**.

---

## 🧱 Clean Architecture Overview

AegisBox backend is built using **Clean Architecture**, ensuring:

- Domain logic is independent
- Infrastructure can be replaced
- HTTP layer is isolated
- Easy unit testing
- Scalable for large features

---

## 🧩 **Architecture Diagram (Mermaid.js)**

```mermaid
flowchart TD

    %% === LAYERS ===
    subgraph Interface Layer
        H[HTTP Router]
        HD[Handlers]
        DTO[DTO Request/Response]
        MW[Middleware]
    end

    subgraph Application Layer
        APP[App Container<br/>Dependency Injection]
    end

    subgraph Domain Layer
        USvc[User Service]
        CSvc[Credential Service]
        ASvc[Audit Service]

        URepoI[User Repository<br/>(Interface)]
        CRepoI[Credential Repository<br/>(Interface)]
        ARepoI[Audit Repository<br/>(Interface)]

        UEntity[User Entity]
        CEntity[Credential Entity]
        AEntity[Audit Entity]
    end

    subgraph Infrastructure Layer
        PGRepoU[User Repository (Postgres)]
        PGRepoC[Credential Repository (Postgres)]
        PGRepoA[Audit Repository (Postgres)]

        DBModelU[User DB Model]
        DBModelC[Credential DB Model]
        DBModelA[Audit DB Model]

        DB[(PostgreSQL)]
    end


    %% === CONNECTIONS ===
    H --> HD
    HD --> DTO
    HD --> APP

    APP --> USvc
    APP --> CSvc
    APP --> ASvc

    USvc --> URepoI
    CSvc --> CRepoI
    ASvc --> ARepoI

    URepoI --> PGRepoU
    CRepoI --> PGRepoC
    ARepoI --> PGRepoA

    PGRepoU --> DBModelU --> DB
    PGRepoC --> DBModelC --> DB
    PGRepoA --> DBModelA --> DB
```

---

## 🔒 Security Features

- **Argon2id password hashing**
- **Per-user salt**
- **Encrypted credential storage**
- **Session management with idle timeout**
- **Login rate limiting + lockout**
- **JSONB audit logging**
- **MFA-ready user model**

---

## 🛠 Running Locally

### ▶ With Docker:

```sh
docker compose up --build
```

### ▶ Without Docker:

```sh
go mod tidy
go run ./cmd/server
```

---
