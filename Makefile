.PHONY: build protobuf clean test run build_client docker_up docker_down docker_build

BUILD_DIR := build
PREFIX := GOOS=linux GOARCH=amd64 CGO_ENABLED=0
build:
	mkdir -p $(BUILD_DIR)
	$(PREFIX) go build -o $(BUILD_DIR)/app_http_server cmd/http_server.go
	$(PREFIX) go build -o $(BUILD_DIR)/app_grpc_server cmd/grpc_server.go

clean:
	rm -fv $(BUILD_DIR)/app*

protobuf:
	protoc --go_out=. --go-grpc_out=. proto/device.proto
	echo "Generate protobuf is done"

run: build
	build/app-devtree-protobuf

test:
	go test ./...

build_client:
	mkdir -p $(BUILD_DIR)
	$(PREFIX) go build -o $(BUILD_DIR)/app_http_client tests/http_client.go
	$(PREFIX) go build -o $(BUILD_DIR)/app_grpc_client tests/grpc_client.go

docker_up:
	docker-compose up -d
docker_down:
	docker-compose down
docker_build:
	docker-compose build