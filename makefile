# Define the binary name
BINARY_NAME=tui-wordle

# Define the build output directory
BUILD_DIR=bin

SRC_DIR=./cmd

# Go commands
GO_BUILD=go build
GO_CLEAN=go clean
GO_TEST=go test
GO_VET=go vet
GO_FMT=gofmt

# Directories
PKG_DIR=./pkg/...
INTERNAL_DIR=./internal/...

# Build for Linux
.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)_linux $(SRC_DIR)

# Build for macOS
.PHONY: build-mac
build-mac:
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)_mac $(SRC_DIR)

# Build for Windows
.PHONY: build-windows
build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)_windows.exe $(SRC_DIR)

# Clean the build output
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)
	mkdir -p $(BUILD_DIR)

# Build all targets
.PHONY: build-all
build-all: clean build-linux build-mac build-windows
