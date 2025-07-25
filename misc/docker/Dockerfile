# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o totp_code_generator main.go

# Final image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/totp_code_generator ./
CMD ["./totp_code_generator"]
