.PHONY: test test-integration all check-test-tags

# Use PowerShell for all recipe lines
SHELL := powershell.exe
.SHELLFLAGS := -NoProfile -Command

# List all packages in the module
PKGS := $(shell go list ./...)

# Run all unit tests (default build, no special tags)
test:
	go clean -testcache
	go test -tags=unit $(PKGS) | Select-String -NotMatch '\[no test files\]'

# Run all integration tests, filtering out "[no test files]"
test-integration:
	go clean -testcache
	go test -tags=integration $(PKGS) | Select-String -NotMatch '\[no test files\]'

# Check that every *_test.go file has a build tag (unit or integration)
check-test-tags:
	& { $$u = Get-ChildItem -Recurse -Include '*_test.go' | Where-Object { -not (Select-String -Quiet -Pattern '^//go:build|^//\+build' -Path $$_.FullName) }; if ($$u.Count -gt 0) { $$u | ForEach-Object { Write-Host "Missing build tag: $$_.Name" }; exit 1 } else { Write-Host 'All test files have build tags.' } }


# Combined: run unit then integration tests
all: test test-integration