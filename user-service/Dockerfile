FROM golang:1.23 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o user-service .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/user-service .

EXPOSE 8082

ENTRYPOINT ["./user-service"]

