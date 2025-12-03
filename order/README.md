example request from grpcurl in windows machine

```sh
grpcurl -plaintext ^
  -d "{\"user_id\":123,\"order_items\":[{\"product_code\":\"prod\",\"quantity\":4,\"unit_price\":12}]}" ^
  localhost:3000 ^
  Order/Create

```
