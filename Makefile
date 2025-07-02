.PHONY: build test lint run docker-up docker-down migrate-up migrate-down clean help

# Variables
APP_NAME=eventflow
CMD_DIR=./cmd/api
BIN_DIR=./bin
MIGRATION_DIR=./migrations
DB_URL=postgres://eventflow_user:eventflow_password@127.0.0.1:5432/eventflow

# Default target
help:
	@echo "Available commands:"
	@echo "  build      - Build the application"
	@echo "  test       - Run tests"
	@echo "  lint       - Run linter"
	@echo "  fmt        - Format code"
	@echo "  pre-commit - Run all quality gates"
	@echo "  run        - Run the application"
	@echo "  docker-up  - Start Docker Compose services"
	@echo "  docker-down- Stop Docker Compose services"
	@echo "  migrate-up - Apply database migrations"
	@echo "  migrate-down - Rollback last migration"
	@echo "  clean      - Clean build artifacts"

# Build application
build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_DIR)/main.go

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run linter
lint:
	@echo "Running linter..."
	golangci-lint run

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Pre-commit hook
pre-commit: fmt lint test build
	@echo "✅ All EventFlow quality gates passed!"


# Run application
run:
	@echo "Running $(APP_NAME)..."
	go run $(CMD_DIR)/main.go

# Docker Compose up
docker-up:
	@echo "Starting Docker Compose services..."
	docker-compose up -d

# Docker Compose down
docker-down:
	@echo "Stopping Docker Compose services..."
	docker-compose down

# Apply database migrations
migrate-up:
	@echo "Applying database migrations..."
	migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" up

# Rollback last migration
migrate-down:
	@echo "Rolling back last migration..."
	migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" down 1

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BIN_DIR)
	go clean
