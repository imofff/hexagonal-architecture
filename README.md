# Hexagonal Go Example 🧱

A modular backend service using Hexagonal Architecture (aka Ports & Adapters) in Go. The project cleanly separates business logic from delivery and infrastructure layers.

## 📦 Tech Stack

- Go 1.24.5
- Gin (HTTP Router)
- GORM (ORM for PostgreSQL)
- PostgreSQL
- Hexagonal Architecture (Ports & Adapters)
- Testify (for unit testing)

## ✅ Features

- Hexagonal Architecture implementation (Entity → Usecase → Adapter)
- PostgreSQL + GORM integration
- Password hashing (bcrypt)
- Environment variable config via `.env`
- SQLite in-memory tests
- Unit and integration tests

## ⚙️ Setup

1. Clone this repo:

```bash
git clone https://github.com/imofff/hexagonal-architecture.git
cd hexagonal-architecture
```

2. Copy `.env.example` to your `.env` and edit as needed:

```bash
cp .env.example .env
```

3. Set up your PostgreSQL database and update `.env` with:

```
DB_HOST=localhost
DB_USER=uruser
DB_PASS=urpassword
DB_NAME=urdbname
DB_PORT=urdbport
PORT=urport
POSTGRES_DB=urdbname
POSTGRES_USER=user
POSTGRES_PASSWORD=pass
```

4. Start PostgreSQL via Docker
```
docker-compose up -d
```
This will run PostgreSQL at localhost:5430.

5. Run the server:

```bash
go run ./cmd/main.go
```

## 🧪 Running Tests

```bash
go test ./... -v
```

Or run specific test file:

```bash
go test ./test/user/user_usecase_test.go
```

## 📁 Folder Structure

```txt
.
├── cmd/                    # Entry point (main.go)
├── config/                 # DB connection, server setup
├── domain/
│   ├── entity/             # Business entity (User)
│   └── port/               # Interface definitions (repository, usecase)
├── internal/
│   ├── adapter/
│   │   ├── http/handler/   # HTTP layer (Gin handlers)
│   │   └── postgres/       # Postgres adapter (GORM)
│   └── app/usecase/        # Business logic implementation
├── test/user/              # Unit & integration tests
├── go.mod / go.sum         # Dependencies
└── .env / .gitignore       # Config & ignores
```

## 🚀 Run

```bash
go run ./cmd/main.go
```


ถ้าอยากได้ `Dockerfile` เพิ่มเพื่อรัน Go app ทั้งโปรเจคใน Docker เลยก็บอกได้นะครับ เดี๋ยวจัดให้ครบทั้ง stack
