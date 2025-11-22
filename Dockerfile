# Build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o /app/server ./cmd/echo_test_server

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /app/server /app/server

ENTRYPOINT ["/app/server"]

EXPOSE 20001
