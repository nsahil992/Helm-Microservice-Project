FROM golang:1.21 AS builder

WORKDIR /app

COPY . .

RUN go build -o user-service .

FROM apline:latest

WORKDIR /app

COPY --from=builder /app/user-service .

EXPOSE 8082

CMD ["./user-service"]

