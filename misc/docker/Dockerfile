# Build stage
FROM golang:1.24-alpine AS builder

RUN echo "Building Docker image"

# Install golangci-lint
RUN apk add --no-cache curl && \
     curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/local/bin v1.59.1

WORKDIR /bin
COPY . .
RUN make build



# Final image
FROM alpine:latest
CMD ["./totp_code_generator"]