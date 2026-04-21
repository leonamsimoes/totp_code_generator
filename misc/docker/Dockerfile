# Build stage
FROM golang:1.26.2-alpine AS builder

RUN echo "Building Docker image"

ENV LINTVERSION=2.50.0
ENV GREMELINSVERSION=0.6.0

# Install golangci-lint
RUN apk add --no-cache curl && \
     curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/local/bin v${LINTVERSION}

# Configuration
RUN apk add --no-cache make

RUN go install github.com/go-gremlins/gremlins/cmd/gremlins@v${GREMELINSVERSION}

WORKDIR /bin
COPY . .
RUN make build

# Final image
FROM alpine:latest
RUN apk add --no-cache make curl
CMD ["./totp_code_generator"]