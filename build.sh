#!/bin/bash

# Application Name
BINARY_NAME="devops-monitor"

echo "Starting cross-platform build pipeline..."

# 1. Compile for Windows (64-bit Intel/AMD)
echo "→ Compiling for Windows (amd64)..."
GOOS=windows GOARCH=amd64 go build -o bin/${BINARY_NAME}-windows-amd64.exe .

# 2. Compile for Linux (64-bit Enterprise/Cloud)
echo "→ Compiling for Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -o bin/${BINARY_NAME}-linux-amd64 .

# 3. Compile for macOS Intel Chips
echo "→ Compiling for macOS (amd64)..."
GOOS=darwin GOARCH=amd64 go build -o bin/${BINARY_NAME}-darwin-amd64 .

# 4. Compile for macOS Apple Silicon Chips (M1/M2/M3)
echo "→ Compiling for macOS (arm64)..."
GOOS=darwin GOARCH=arm64 go build -o bin/${BINARY_NAME}-darwin-arm64 .

echo "✅ Build pipeline complete! Check your /bin folder."