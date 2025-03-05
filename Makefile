build:
	@go build -o bin/backend cmd/backend/main.go

run: build
	@./bin/backend

test:
	@go test -v ./...