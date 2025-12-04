#!/bin/bash

# Path to Go services
GW_SVC=../last-legends-gateway-service
# USER_SVC=../user-service
# ORDER_SVC=../order-service

set -e

GATEWAY_PATH=$(go list -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway/v2)

if [ ! -d "third_party/googleapis" ]; then
  echo "Google APIs not found. Cloning..."
  mkdir -p third_party
  git clone https://github.com/googleapis/googleapis.git third_party/googleapis
fi

# ensure plugins installed
for plugin in protoc-gen-go protoc-gen-go-grpc protoc-gen-grpc-gateway protoc-gen-openapiv2; do
  if ! command -v $plugin &>/dev/null; then
    echo "$plugin not found, installing..."
    case $plugin in
      protoc-gen-go) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest ;;
      protoc-gen-go-grpc) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest ;;
      protoc-gen-grpc-gateway) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest ;;
      protoc-gen-openapiv2) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest ;;
    esac
  fi
done

# User-service & swagger
protoc -I . \
  -I third_party/googleapis \
  -I $GATEWAY_PATH \
  --go_out user-proto/pb --go-grpc_out user-proto/pb \
  --grpc-gateway_out user-proto/pb \
  --openapiv2_out . \
  user-proto/proto/user.proto

mv user-proto/proto/user.swagger.json $GW_SVC/docs/swagger/user.swagger.json

# Product-service & swagger
protoc -I . \
  -I third_party/googleapis \
  -I $GATEWAY_PATH \
  --go_out product-proto/pb --go-grpc_out product-proto/pb \
  --grpc-gateway_out product-proto/pb \
  --openapiv2_out . \
  product-proto/proto/product.proto

mv product-proto/proto/product.swagger.json $GW_SVC/docs/swagger/product.swagger.json