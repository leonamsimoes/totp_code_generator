# Build stage
FROM golang:1.24-alpine AS builder


RUN apk update
RUN apk add make # install make

WORKDIR /app
COPY . .
RUN make build

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.59.1

WORKDIR /bin
COPY /bin/ ./
CMD ["./totp_code_generator"]