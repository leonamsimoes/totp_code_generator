# Build stage
FROM golang:1.24-alpine AS builder

# Install golangci-lint
RUN apk add --no-cache curl && \
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/local/bin v1.59.1

WORKDIR /app
COPY . .
RUN go build -o totp_code_generator main.go

# Final image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/totp_code_generator ./
CMD ["./totp_code_generator"]
