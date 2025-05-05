FROM golang:1.21 AS builder

WORKDIR /app

COPY . .

RUN go build -o greet-service .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/greet-service .

EXPOSE 8080

CMD ["./greet-service"]