FROM golang:1.23 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o greet-service .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/greet-service .

EXPOSE 8080

ENTRYPOINT ["./greet-service"]