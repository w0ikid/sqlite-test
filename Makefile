include .env

# postgres migrate and rollback

migrate-postgres:
	@echo "Running PostgreSQL migrations..."
	migrate -path migrations/postgres -database "postgres://$(PG_USER):$(PG_PASSWORD)@$(PG_HOST):$(PG_PORT)/$(PG_DB)?sslmode=${PG_SSLMODE}" up

rollback-postgres:
	@echo "Rolling back PostgreSQL migrations..."
	migrate -path migrations/postgres -database "postgres://$(PG_USER):$(PG_PASSWORD)@$(PG_HOST):$(PG_PORT)/$(PG_DB)?sslmode=${PG_SSLMODE}" down 1


# sqlite migrate and rollback

migrate-sqlite:
	@echo "Running SQLite migrations..."
	migrate -path migrations/sqlite -database "sqlite://$(SQLITE_PATH)" up

rollback-sqlite:
	@echo "Rolling back SQLite migrations..."
	migrate -path migrations/sqlite -database "sqlite://$(SQLITE_PATH)" down 1

run:
	@echo "Running application..."
	@go run cmd/app/main.go