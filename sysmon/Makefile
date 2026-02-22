build:
	@go build -o bin/sysmon cmd/sysmon/main.go

test:
	@go test -v ./...

run: build
	@./bin/sysmon

