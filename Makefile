.PHONY: test test-integration all

# Use PowerShell for all recipe lines
SHELL := powershell.exe
.SHELLFLAGS := -NoProfile -Command

# List all packages in the module
PKGS := $(shell go list ./...)

# Run all unit tests (default build, no special tags)
test:
	go clean -testcache
	go test $(PKGS)

# Run all integration tests, filtering out "[no test files]"
test-integration:
	go clean -testcache
	go test -tags=integration $(PKGS) | Select-String -NotMatch '\[no test files\]'

# Combined: run unit then integration tests
all: test test-integration