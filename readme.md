### microservices-order

run database in local

```sh
docker run -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=verysecretpass \
    -e MYSQL_DATABASE=order mysql

```

run the order service application

```sh
DATA_SOURCE_URL=root:verysecretpass@tcp(127.0.0.1:3306)/order \
APPLICATION_PORT=3000 \
ENV=development \
go run cmd/main.go
```

calling gRPC endpoint with `grpcurl`

```sh
grpcurl -d '{"user_id": 123, "order_items": [{"product_code": "prod", "quantity": 4, "unit_price": 12}]}' -plaintext localhost:3000 Order/Create
```
