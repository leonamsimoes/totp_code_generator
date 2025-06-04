PROJECT_NAME = totp_code_generator

LDFLAGS_PATH = gihub.com/totp_code_generator
LDFLAGS ?= "-X ${LDFLAGS_PATH}runtime.version=$(shell git rev-parse --short HEAD)"

test: # Run Go tests after starting services
	@echo "Running Go tests..."
	--env local --token=$(TOKEN)
	go clean -cache
	go test ./... -cover -timeout 5m

lint: # Linter
	@echo "=== Lint code ==="
	golangci-lint --version
	golangci-lint run