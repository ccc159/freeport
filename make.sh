#!/bin/bash

# Name of the output binary
BINARY_NAME="freeport"

# Directories to store the builds
BUILD_DIR="build"
LINUX_DIR="${BUILD_DIR}/linux"
WINDOWS_DIR="${BUILD_DIR}/windows"
MAC_DIR="${BUILD_DIR}/mac"

# Clear existing build directories
echo "Clearing existing builds..."
rm -rf "${BUILD_DIR}"

# Create the build directories
mkdir -p "${LINUX_DIR}" "${WINDOWS_DIR}" "${MAC_DIR}"

# Initialize Go module if go.mod doesn't exist
if [ ! -f go.mod ]; then
    echo "Initializing Go module..."
    go mod init "${BINARY_NAME}"
fi

echo "Tidying up Go modules..."
go mod tidy

echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o "${LINUX_DIR}/${BINARY_NAME}"

echo "Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o "${WINDOWS_DIR}/${BINARY_NAME}.exe"

echo "Building for macOS..."
GOOS=darwin GOARCH=amd64 go build -o "${MAC_DIR}/${BINARY_NAME}"

echo "Builds completed successfully:"
echo "- ${LINUX_DIR}/${BINARY_NAME}"
echo "- ${WINDOWS_DIR}/${BINARY_NAME}.exe"
echo "- ${MAC_DIR}/${BINARY_NAME}"
