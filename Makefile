# Go Bookstore Application Makefile

# Variables
APP_NAME=bookstore
MAIN_PATH=cmd/main
BUILD_DIR=bin
DOCKER_IMAGE=bookstore-app

# Go related variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=$(BUILD_DIR)/$(APP_NAME)
BINARY_UNIX=$(BUILD_DIR)/$(APP_NAME)_unix

# Default target
.DEFAULT_GOAL := help

# Help target
.PHONY: help
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Build targets
.PHONY: build
build: ## Build the application
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_PATH)

.PHONY: build-linux
build-linux: ## Build for Linux
	@echo "Building for Linux..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) $(MAIN_PATH)

.PHONY: build-windows
build-windows: ## Build for Windows
	@echo "Building for Windows..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME).exe $(MAIN_PATH)

.PHONY: build-mac
build-mac: ## Build for macOS
	@echo "Building for macOS..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)_mac $(MAIN_PATH)

# Run targets
.PHONY: run
run: ## Run the application
	@echo "Running $(APP_NAME)..."
	$(GOCMD) run $(MAIN_PATH)/main.go

.PHONY: run-dev
run-dev: ## Run with development environment
	@echo "Running in development mode..."
	APP_ENV=development $(GOCMD) run $(MAIN_PATH)/main.go

.PHONY: run-prod
run-prod: ## Run with production environment
	@echo "Running in production mode..."
	APP_ENV=production $(GOCMD) run $(MAIN_PATH)/main.go

# Test targets
.PHONY: test
test: ## Run tests
	@echo "Running tests..."
	$(GOTEST) -v ./...

.PHONY: test-coverage
test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

.PHONY: test-short
test-short: ## Run short tests
	@echo "Running short tests..."
	$(GOTEST) -v -short ./...

# Clean targets
.PHONY: clean
clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	$(GOCLEAN)
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html

# Dependency management
.PHONY: deps
deps: ## Download dependencies
	@echo "Downloading dependencies..."
	$(GOGET) -v -t -d ./...

.PHONY: deps-update
deps-update: ## Update dependencies
	@echo "Updating dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

.PHONY: deps-vendor
deps-vendor: ## Vendor dependencies
	@echo "Vendoring dependencies..."
	$(GOMOD) vendor

# Linting and formatting
.PHONY: fmt
fmt: ## Format code
	@echo "Formatting code..."
	$(GOCMD) fmt ./...

.PHONY: vet
vet: ## Vet code
	@echo "Vetting code..."
	$(GOCMD) vet ./...

.PHONY: lint
lint: ## Run linter (requires golangci-lint)
	@echo "Running linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not found. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# Swagger documentation
.PHONY: swagger-init
swagger-init: ## Generate Swagger documentation
	@echo "Generating Swagger documentation..."
	@if command -v swag >/dev/null 2>&1; then \
		swag init -g $(MAIN_PATH)/main.go -o ./docs; \
		echo "Swagger documentation generated in ./docs/"; \
	else \
		echo "swag not found. Install with: go install github.com/swaggo/swag/cmd/swag@latest"; \
	fi

.PHONY: swagger-serve
swagger-serve: ## Serve Swagger documentation (requires swagger-ui)
	@echo "Serving Swagger documentation..."
	@if command -v swagger >/dev/null 2>&1; then \
		swagger serve -F swagger docs/swagger.json; \
	else \
		echo "swagger not found. Install with: go install github.com/go-swagger/go-swagger/cmd/swagger@latest"; \
	fi

# Docker targets
.PHONY: docker-build
docker-build: ## Build Docker image
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

.PHONY: docker-run
docker-run: ## Run Docker container
	@echo "Running Docker container..."
	docker run -p 8080:8080 --env-file .env $(DOCKER_IMAGE)

.PHONY: docker-clean
docker-clean: ## Clean Docker images
	@echo "Cleaning Docker images..."
	docker rmi $(DOCKER_IMAGE) 2>/dev/null || true

# Development setup
.PHONY: setup
setup: ## Setup development environment
	@echo "Setting up development environment..."
	@if [ ! -f .env ]; then \
		echo "Creating .env file from config.example..."; \
		cp config.example .env; \
		echo "Please update .env with your actual credentials"; \
	else \
		echo ".env file already exists"; \
	fi
	$(GOMOD) download
	@echo "Development environment setup complete!"

.PHONY: install-tools
install-tools: ## Install development tools
	@echo "Installing development tools..."
	$(GOCMD) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(GOCMD) install github.com/joho/godotenv/cmd/godotenv@latest
	$(GOCMD) install github.com/swaggo/swag/cmd/swag@latest
	@echo "Development tools installed!"

# Database targets
.PHONY: db-migrate
db-migrate: ## Run database migrations (placeholder)
	@echo "Database migrations not implemented yet"
	@echo "Add your migration logic here"

.PHONY: db-seed
db-seed: ## Seed database with sample data (placeholder)
	@echo "Database seeding not implemented yet"
	@echo "Add your seeding logic here"

# Security targets
.PHONY: security-check
security-check: ## Run security checks
	@echo "Running security checks..."
	@if command -v gosec >/dev/null 2>&1; then \
		gosec ./...; \
	else \
		echo "gosec not found. Install with: go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest"; \
	fi

# Performance targets
.PHONY: bench
bench: ## Run benchmarks
	@echo "Running benchmarks..."
	$(GOCMD) test -bench=. ./...

.PHONY: profile
profile: ## Generate CPU profile
	@echo "Generating CPU profile..."
	$(GOCMD) test -cpuprofile=cpu.prof -bench=. ./...

# Release targets
.PHONY: release
release: ## Build release binaries
	@echo "Building release binaries..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME)-linux-amd64 $(MAIN_PATH)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME)-windows-amd64.exe $(MAIN_PATH)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME)-darwin-amd64 $(MAIN_PATH)
	@echo "Release binaries created in $(BUILD_DIR)/"

# Utility targets
.PHONY: version
version: ## Show version information
	@echo "Application: $(APP_NAME)"
	@echo "Go version: $(shell go version)"
	@echo "Build time: $(shell date)"

.PHONY: check
check: ## Run all checks (fmt, vet, test)
	@echo "Running all checks..."
	$(MAKE) fmt
	$(MAKE) vet
	$(MAKE) test

.PHONY: all
all: clean deps test build ## Clean, get deps, test, and build 