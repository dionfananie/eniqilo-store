# BeliMang Backend Service

This is a backend service that provides api for store

## Getting Started

### Prerequisites

- Go 1.22
- Postgres
- Golang Migrate CLI
- Docker

### Installation

1. Migrate the database schema

```sh
migrate -database "postgres://postgres:password@localhost:5432/eniqilodb?sslmode=disable" -path ./db/migrations -verbose up
```

2. Run the application

```sh
go run main.go
```

Specification App
https://www.notion.so/dionfananie/EniQilo-Store-e433f5c9cbb64c638c2afefae930e999

Test K6
https://github.com/dionfananie/EniQiloStoreTestCasesPSW2B2
