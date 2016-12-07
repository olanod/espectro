default: cover

BIN=bin/espectro

test:
	@go test -coverprofile=cover.out

cover: test
	@echo "Coverage info:"
	@go tool cover -func=cover.out

build:
	@go build -o $(BIN) cli/main.go
