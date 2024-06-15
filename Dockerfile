# Start with the official Golang image for building the application
FROM golang:1.21.5 AS builder
RUN apt-get update && apt-get install -y libudev-dev && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build

FROM ubuntu:latest
RUN apt-get update && apt-get install -y ca-certificates libudev-dev && rm -rf /var/lib/apt/lists/*
WORKDIR /root/
COPY --from=builder /app/build/app_http_server .
COPY --from=builder /app/build/app_grpc_server .

EXPOSE 8080
CMD ["./app_grpc_server"]
