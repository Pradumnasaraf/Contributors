FROM golang:1.19.5-alpine3.17 AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 8080
CMD ["go", "run", "main.go"]
