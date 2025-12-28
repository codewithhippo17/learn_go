# Agent Guidelines for Go Project

This document provides essential guidelines for AI coding agents working on this Go codebase.

## Project Overview

- **Language**: Go 1.25.5
- **Module**: example/hello
- **Type**: Simple Go application (currently a "Hello World" example)
- **Build System**: Standard Go toolchain
- **Dependencies**: None (standard library only)

## Build, Test, and Lint Commands

### Build Commands
```bash
# Build the application
go build

# Build with specific output name
go build -o hello

# Build for different platforms
GOOS=linux GOARCH=amd64 go build -o hello-linux
GOOS=windows GOARCH=amd64 go build -o hello.exe

# Run without building
go run .
go run hello.go
```

### Test Commands
```bash
# Run all tests
go test

# Run tests with verbose output
go test -v

# Run tests in all subdirectories
go test ./...

# Run a specific test
go test -run TestFunctionName

# Run tests with race detection
go test -race

# Run tests with coverage
go test -cover
go test -coverprofile=coverage.out

# View coverage in browser
go tool cover -html=coverage.out
```

### Lint and Format Commands
```bash
# Format code (modifies files in-place)
go fmt ./...

# Check formatting without modifying
gofmt -l .

# Run Go vet (static analysis)
go vet ./...

# Run additional linters (if golangci-lint is installed)
golangci-lint run

# Fix imports
goimports -w .
```

### Module Management
```bash
# Add a dependency
go get github.com/example/package

# Remove unused dependencies
go mod tidy

# Update dependencies
go get -u ./...

# Verify dependencies
go mod verify
```

## Code Style Guidelines

### Imports
- Use standard library imports first, then third-party, then local packages
- Group imports with blank lines between groups
- Use `goimports` to automatically organize imports

```go
package main

import (
    "context"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/sirupsen/logrus"

    "example/hello/internal/config"
    "example/hello/pkg/utils"
)
```

### Formatting
- Use `gofmt` for consistent formatting
- Use tabs for indentation (Go standard)
- Line length: aim for 80-100 characters, but readability takes precedence
- Use blank lines to separate logical sections

### Types and Declarations
```go
// Prefer explicit type declarations for clarity
var count int = 0

// Use short variable declarations when type is obvious
name := "example"

// Group related constants
const (
    MaxRetries = 3
    Timeout    = 30 * time.Second
)

// Group related variables
var (
    ErrNotFound = errors.New("not found")
    ErrInvalid  = errors.New("invalid input")
)
```

### Naming Conventions
- **Packages**: lowercase, single word when possible (`http`, `json`)
- **Functions/Methods**: CamelCase, exported functions start with uppercase
- **Variables**: camelCase, exported variables start with uppercase
- **Constants**: CamelCase, exported constants start with uppercase
- **Interfaces**: Usually end with `-er` (`Reader`, `Writer`, `Handler`)

```go
// Good
func ProcessUserData(user User) error
func (u *User) GetFullName() string
var userCount int
const DefaultTimeout = 30

// Bad
func process_user_data(user User) error
func (u *User) get_full_name() string
var user_count int
const default_timeout = 30
```

### Error Handling
- Always handle errors explicitly
- Use `fmt.Errorf` with `%w` verb for error wrapping
- Return errors as the last return value
- Use sentinel errors for expected error conditions

```go
// Good error handling
func processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file %s: %w", filename, err)
    }
    defer file.Close()

    // Process file...
    if err := someOperation(); err != nil {
        return fmt.Errorf("operation failed: %w", err)
    }
    
    return nil
}
```

### Function Design
- Keep functions small and focused on a single responsibility
- Use meaningful parameter and return names
- Prefer explicit error returns over panics
- Use context.Context for cancellation and timeouts

```go
func fetchUserData(ctx context.Context, userID string) (*User, error) {
    // Implementation
}
```

### Comments and Documentation
- Use `godoc` style comments for exported functions, types, and packages
- Start comments with the name being documented
- Write complete sentences

```go
// ProcessUser validates and processes user data.
// It returns an error if the user data is invalid.
func ProcessUser(user *User) error {
    // Implementation
}
```

### Testing Guidelines
- Place tests in `*_test.go` files
- Use table-driven tests for multiple test cases
- Test both positive and negative scenarios
- Use `testing.T` for unit tests, `testing.B` for benchmarks

```go
func TestProcessUser(t *testing.T) {
    tests := []struct {
        name    string
        user    *User
        wantErr bool
    }{
        {"valid user", &User{Name: "John"}, false},
        {"empty name", &User{Name: ""}, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := ProcessUser(tt.user)
            if (err != nil) != tt.wantErr {
                t.Errorf("ProcessUser() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

## Project Structure Guidelines

As the project grows, follow this recommended structure:

```
.
├── cmd/                    # Main applications
│   └── hello/
│       └── main.go
├── internal/               # Private application code
│   ├── config/
│   ├── handler/
│   └── service/
├── pkg/                    # Library code (can be imported by others)
│   └── utils/
├── api/                    # API definitions (OpenAPI, protobuf)
├── docs/                   # Documentation
├── scripts/                # Build and deployment scripts
├── test/                   # Integration tests
├── go.mod
├── go.sum
├── Makefile               # Build automation
└── README.md
```

## Additional Notes

- This project currently has no external dependencies
- No Cursor rules or Copilot instructions found
- Consider adding `golangci-lint` for additional static analysis
- Consider adding a `Makefile` for build automation
- Consider adding CI/CD configuration (GitHub Actions, etc.)
- Follow Go proverbs: "Don't communicate by sharing memory, share memory by communicating"

## Useful Commands for Agents

```bash
# Quick project health check
go mod tidy && go vet ./... && go test ./...

# Full code quality check (if golangci-lint available)
golangci-lint run && go test -race -cover ./...

# Generate documentation
godoc -http=:6060
```

Remember to always run `go fmt` and `go vet` before committing code changes.