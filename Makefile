.PHONY: build clean test

# Binary name
BINARY_NAME=assembly_time
# Build directory
BIN_DIR=bin

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(BINARY_NAME) .

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(BIN_DIR)

# Run tests
test:
	@go test -v ./...

# Default target
all: build

# Install binary
install: build
	@echo "Installing $(BINARY_NAME)..."
	@cp $(BIN_DIR)/$(BINARY_NAME) ${GOPATH}/bin/$(BINARY_NAME) 