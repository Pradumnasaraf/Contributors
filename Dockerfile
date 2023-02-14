FROM golang:1.19.5-alpine3.17 AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app

FROM alpine:3.14.2
COPY --from=builder /app /bin/app
ENTRYPOINT ["/bin/app"]
