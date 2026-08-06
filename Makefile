# Makefile for Viva Las Mesh

.PHONY: build test vet lint stats clean install

# Build targets
build:
	go build -o ./bin/viva-las-mesh ./cmd/mesh

build-daemon:
	go build -o ./bin/orchestrator-server ./cmd/orchestrator-server

# Cross-platform build targets
cross-build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/viva-las-mesh-linux ./cmd/mesh

cross-build-macos:
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/viva-las-mesh-darwin ./cmd/mesh

cross-build-windows:
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/viva-las-mesh-windows.exe ./cmd/mesh

# Test targets
test:
	go test -race ./...

test-verbose:
	go test -race -v ./...

test-cover:
	go test -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Lint & Vet targets
vet:
	go vet ./...
lint:
	golangci-lint run ./...

# Stats analysis
stats:
	go-stats-generator analyze . --skip-tests --format json --sections functions,duplication,documentation -o tmp/stats.json
stats-diff:
	go-stats-generator diff tmp/baseline-exec.json tmp/stats.json

# Install tools
install-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/opd-ai/go-stats-generator@latest

# Clean
clean:
	rm -rf ./bin ./tmp coverage.out coverage.html

# Default target
all: build test vet
