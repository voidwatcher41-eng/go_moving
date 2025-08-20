# Go Moving

This repository contains a small educational program that demonstrates modern
Go practices such as context-aware HTTP requests, structured error handling and
unit testing. The application fetches a Todo item from the JSONPlaceholder API,
modifies it and prints the original and updated versions.

## Getting started

### Prerequisites
- [Go 1.24+](https://go.dev/dl/)

### Install dependencies
```bash
go mod tidy
```

### Run the program
```bash
go run .
```

## Development

The project follows common Go development workflows:

```bash
# Format source files
go fmt ./...

# Run static analysis
go vet ./...

# Execute tests
go test ./...
```

If [golangci-lint](https://golangci-lint.run/) is available, running
`golangci-lint run` will provide extended linting.

## Project structure

```
main.go       // Program entry point and high-level orchestration
todo.go       // Todo type and FetchTodo helper
todo_test.go  // Unit tests for FetchTodo
```

## Learning notes

- **Context and timeouts:** `FetchTodo` accepts a `context.Context` and the
  main function uses `context.WithTimeout` and an HTTP client with a timeout to
  avoid hanging network calls.
- **Structured errors:** Errors are wrapped with additional context using
  `fmt.Errorf` to make troubleshooting easier.
- **Testing with httptest:** The tests spin up lightweight HTTP servers to
  simulate API responses, demonstrating effective unit testing of HTTP
  clients.

