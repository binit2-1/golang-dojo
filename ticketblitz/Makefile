# Define the DB URL once as a variable
DB_URL := "postgresql://ticketblitz-admin:password123@localhost:5432/ticketblitz-db?sslmode=disable"

.PHONY: build test run db-it migrate-up migrate-down migrate-force migrate-drop

build:
	@go build -o bin/ticketblitz cmd/api/main.go

test:
	@go test -v ./...

run: build
	@./bin/ticketblitz

db-it: 
	@docker exec -it ticketblitz-db psql -U ticketblitz-admin -d ticketblitz-db

migrate-up:
	@migrate -path migrations -database $(DB_URL) up

migrate-down:
	@migrate -path migrations -database $(DB_URL) down

# FIX: Forces the migration version to fix the "dirty database" error. 
# Usage: make migrate-force V=0
migrate-force:
	@migrate -path migrations -database $(DB_URL) force $(V)

# NUCLEAR OPTION: Drops everything inside the database (all tables and migration history)
# Highly recommended for local development when you completely mess up.
migrate-drop:
	@migrate -path migrations -database $(DB_URL) drop -f