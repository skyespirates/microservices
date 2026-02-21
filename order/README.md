example request from grpcurl in windows machine

```sh
grpcurl -plaintext ^
  -d "{\"user_id\":123,\"order_items\":[{\"product_code\":\"prod\",\"quantity\":4,\"unit_price\":12}]}" ^
  localhost:3000 ^
  Order/Create

```

example request for linux

```sh
grpcurl -plaintext \
  -d '{"user_id":112,"order_items":[{"product_code":"aym","quantity":3,"unit_price":8}]}' \
  localhost:3000 Order/Create
```

```sh
grpcurl -plaintext \
  -d '{"order_id":1}' \
  localhost:3000 Order/Get
```
