# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

This project uses `just` as a task runner (replacement for make). All commands are defined in the `justfile`.

### Essential Commands
- `just qa` - Complete quality assurance: go mod tidy/verify, vet, format with golangci-lint, lint, and vulnerability check
- `just unittest` - Run tests with coverage report (generates coverage.out and coverage.html)  
- `just benchmark` - Run benchmarks and compare with previous results using benchstat
- `just doc` - Start documentation server at localhost:8080

### Development Workflow
After rebasing or initial setup:
```bash
just tools  # Install required Go tools (referenced in DEVELOPMENT.md)
just qa     # Quality check all code
```

When changing code:
```bash
just qa     # Always run before committing
```

## Code Architecture

This is a Go library providing miscellaneous list utility functions. The codebase is intentionally minimal:

- `list.go` - Main library with generic functions using Go generics (T comparable)
- `list_test.go` - Comprehensive example-based tests using Go's example testing pattern
- `benchmark_test.go` - Performance benchmarks

### Key Implementation Details
- Uses Go 1.24 features including generics with comparable constraint
- Functions operate on generic slices `[]T` where `T` is comparable
- Primary function `Conflate[T comparable](slice []T, limit int) []T` removes consecutive duplicates exceeding a specified limit
- Helper function `countConsecutive` counts consecutive identical elements
- Tests use Go's example testing convention with `ExampleFunction_scenario` naming

### Testing Approach
- Uses Go's built-in example testing (functions starting with `Example`)
- Output verification through `// Output:` comments
- Coverage reporting enabled via `go test -coverprofile`
- Benchmarking with memory allocation tracking (`-benchmem`)