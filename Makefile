include .env

MIGRATE_BIN := migrate
SQLITE_URL := sqlite://$(SQLITE_PATH)?x-migrations-table=schema_migrations
POSTGRES_URL := postgres://$(PG_USER):$(PG_PASSWORD)@$(PG_HOST):$(PG_PORT)/$(PG_DB)?sslmode=${PG_SSLMODE}

# postgres migrate and rollback

migrate-postgres:
	@echo "Running PostgreSQL migrations..."
	$(MIGRATE_BIN) -path migrations/postgres -database "$(POSTGRES_URL)" up

rollback-postgres:
	@echo "Rolling back PostgreSQL migrations..."
	$(MIGRATE_BIN) -path migrations/postgres -database "$(POSTGRES_URL)" down 1


# sqlite migrate and rollback

migrate-sqlite:
	@echo "Running SQLite migrations..."
	$(MIGRATE_BIN) -path migrations/sqlite -database "$(SQLITE_URL)" up

rollback-sqlite:
	@echo "Rolling back SQLite migrations..."
	$(MIGRATE_BIN) -path migrations/sqlite -database "$(SQLITE_URL)" down 1

# Create a new SQLite migration: make migrate-create-sqlite name=create_users
migrate-create-sqlite:
ifndef name
	$(error name is required. Use: make migrate-create-sqlite name=create_users)
endif
	$(MIGRATE_BIN) create -ext sql -dir migrations/sqlite -seq $(name)


run:
	@echo "Running application..."
	@go run cmd/app/main.go