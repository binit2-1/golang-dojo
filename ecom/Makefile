build:
	@go build -o bin/ecom cmd/main.go

test:
	@go test -v ./..

run: build
	@./bin/ecom

db-up:
	@docker-compose -f db.docker-compose.yaml up -d

db-down:
	@docker-compose -f db.docker-compose.yaml down