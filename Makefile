PROJECT_NAME 		   = totp_code_generator
LDFLAGS_PATH 		   = gihub.com/
PROJECT 	 		   = gihub.com/totp_code_generator
CI_IMAGE_VERSION	   = 001
DOCKER_REPOSITORY	   = leonamsimoes/projects
FULL_DOCKER_REPOSITORY = ${DOCKER_REPOSITORY}:${PROJECT_NAME}_${CI_IMAGE_VERSION}

 
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

docker:
	@echo "<.:: Docker Image ::.>"
	docker pull ${FULL_DOCKER_REPOSITORY}
	docker build --tag ${FULL_DOCKER_REPOSITORY} -f misc/docker/Dockerfile .
	docker login
	docker push ${FULL_DOCKER_REPOSITORY}