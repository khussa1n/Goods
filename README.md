# Goods Application

By Khussain Kudaibergenov | Хусаин Кудайбергенов

## How to Run

1. Make sure you have Docker and Docker Compose installed on your system.

2. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/goods-app.git

3. Run
    ```bash
   docker-compose up --build

This repository contains a Docker Compose setup for the Goods application, which consists of several services: PostgreSQL, Redis, ClickHouse, NATS, Gin server, and Golang Server which connected to clickhouse.

## Services

### PostgreSQL
- Main DB
- Port: 5437 (host) mapped to 5432 (container)

### Redis
- DB for cache
- Port: 6380 (host) mapped to 6379 (container)

### ClickHouse
- DB for data logs
- Port: 9001 (host) mapped to 9000 (container)

### NATS
- Message broker between Gin server and ClickHouse
- Port: 4223 (host) mapped to 4222 (container)

### Gin server with PostgreSQL and  Redis
- Build: ./app_sender
- Port: 8081 (host) mapped to 8080 (container)

### Golang server with ClickHouse
- Build: ./app_receiver
- Port: 8082 (host) mapped to 8082 (container)