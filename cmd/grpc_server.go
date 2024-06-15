package main

import (
	"github.com/elephant315/devtree-protobuf/internal/adapters/device"
	gr "github.com/elephant315/devtree-protobuf/internal/adapters/grpc"
	"github.com/elephant315/devtree-protobuf/internal/ports"
	"github.com/elephant315/devtree-protobuf/internal/service"
	pb "github.com/elephant315/devtree-protobuf/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listn, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to start listen: %v", err)
	}

	var deviceAdapter ports.DeviceAdapter = device.NewDeviceAdapter()
	deviceService := service.NewDeviceService(deviceAdapter)
	grpcServer := gr.NewGRPCServer(deviceService)

	s := grpc.NewServer()
	pb.RegisterDeviceServiceServer(s, grpcServer)

	log.Println("gRPC server is running on port 50051")
	if err := s.Serve(listn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
