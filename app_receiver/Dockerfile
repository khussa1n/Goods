FROM golang:1.22.0-alpine3.19 AS builder

WORKDIR /app
COPY .. .
RUN go build -o receiver ./cmd/main.go

FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/receiver .
COPY config.yaml .
COPY migrations/000001_init.up.sql /app/migrations/000001_init.up.sql
COPY migrations/000001_init.down.sql /app/migrations/000001_init.down.sql

EXPOSE 8082

CMD ["./receiver"]