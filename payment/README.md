```sh
grpcurl -plaintext \
  -d '{"user_id": 1, "order_id":1, "total_price": 3.14}' \
  localhost:3001 Payment/Create
```
