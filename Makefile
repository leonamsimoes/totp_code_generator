PROJECT_NAME = totp_code_generator

LDFLAGS_PATH = gihub.com/totp_code_generator
LDFLAGS ?= "-X ${LDFLAGS_PATH}runtime.version=$(shell git rev-parse --short HEAD)"

install: version # Install vendor
	@echo "<.:: Running Go tests ::.>"
	go mod vendor


test: # Run Go tests after starting services
	@echo "<.:: Running Go tests ::.>"
	go test ./... -cover -timeout 5m


lint: # Linter
	@echo "<.:: Lint code ::.>"
	golangci-lint --version ;
	golangci-lint run


quality: test lint
	@echo "<.:: Quality code ::.>"


build: # Build the binary 
	@echo "<.:: Building ::.>"
	go build -o bin/totp_code_generator main.go


version: # Printing the version
	@echo "<.:: Versions ::.>"
	go version ;
	golangci-lint --version ;