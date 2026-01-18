#!/bin/bash

# Golang Interview Handbook - Test Runner Script
# This script runs all tests with proper formatting and output

set -e

echo "========================================="
echo "Golang Interview Handbook - Test Runner"
echo "========================================="
echo ""

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Run all tests
echo "Running all tests..."
echo ""

if go test ./... -v; then
    echo ""
    echo -e "${GREEN}✓ All tests passed!${NC}"
    echo ""
    
    # Show test coverage
    echo "Generating coverage report..."
    go test ./... -cover
    echo ""
    
    # Run with race detector
    echo "Running with race detector..."
    if go test ./... -race; then
        echo -e "${GREEN}✓ No race conditions detected!${NC}"
    else
        echo -e "${RED}✗ Race conditions detected!${NC}"
        exit 1
    fi
    
    echo ""
    echo "Running go vet..."
    if go vet ./...; then
        echo -e "${GREEN}✓ go vet passed!${NC}"
    else
        echo -e "${RED}✗ go vet found issues!${NC}"
        exit 1
    fi
    
    echo ""
    echo -e "${GREEN}========================================="
    echo "All checks passed! 🎉"
    echo "=========================================${NC}"
else
    echo ""
    echo -e "${RED}✗ Tests failed!${NC}"
    exit 1
fi
