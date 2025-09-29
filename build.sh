#!/bin/bash

# Path to Go services
USER_SVC=../user-service
# ORDER_SVC=../order-service

# User-service
protoc -I . \
  --go_out $USER_SVC/pb --go-grpc_out $USER_SVC/pb \
  --grpc-gateway_out $USER_SVC/pb \
  --openapiv2_out $USER_SVC/docs \
  user-proto/proto/user.proto

# Order-service
# protoc -I . \
#   --go_out $ORDER_SVC/internal/pb --go-grpc_out $ORDER_SVC/internal/pb \
#   --grpc-gateway_out $ORDER_SVC/internal/pb \
#   --openapiv2_out $ORDER_SVC/docs \
#   order/order.proto
