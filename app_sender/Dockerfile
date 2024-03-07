FROM golang:1.22.0-alpine3.19 AS builder

WORKDIR /app
COPY .. .
RUN go build -o sender ./cmd/main.go

FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/sender .
COPY config.yaml .
COPY migrations/pg/000001_init.up.sql /app/migrations/pg/000001_init.up.sql
COPY migrations/pg/000001_init.down.sql /app/migrations/pg/000001_init.down.sql

EXPOSE 8080
RUN apk update && apk add postgresql-client

CMD ["./sender"]