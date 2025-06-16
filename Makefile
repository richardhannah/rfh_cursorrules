.PHONY: test test-integration

# List all packages in the module
PKGS := $(shell go list ./...)

# Run all unit tests (default build, no special tags)
test:
	go test $(PKGS)

# Run all integration tests (those guarded by // +build integration)
# and any other tests marked with the 'integration' tag
test-integration:
	go test -tags=integration $(PKGS)

# Combined: run unit then integration tests
all: test test-integration