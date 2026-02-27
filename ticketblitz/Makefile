build:
	@go build -o bin/ticketblitz cmd/api/main.go

test:
	@go test -v ./...

run: build
	@./bin/ticketblitz

db-it: 
	@docker exec -it ticketblitz-db psql -U ticketblitz-admin -d ticketblitz-db

migrate-up:
	@migrate -path migrations -database "postgresql://ticketblitz-admin:password123@localhost:5432/ticketblitz-db?sslmode=disable" up
