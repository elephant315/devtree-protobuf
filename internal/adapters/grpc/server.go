package grpc

import (
	"context"
	"github.com/elephant315/devtree-protobuf/internal/ports"
	pb "github.com/elephant315/devtree-protobuf/proto"

	"google.golang.org/grpc/status"
)

type GRPCServer struct {
	pb.UnimplementedDeviceServiceServer
	service ports.DeviceService
}

func NewGRPCServer(service ports.DeviceService) *GRPCServer {
	return &GRPCServer{
		service: service,
	}
}

func (s *GRPCServer) GetDevices(ctx context.Context, req *pb.GetDevicesRequest) (*pb.DeviceList, error) {
	devices, err := s.service.GetDevices()
	if err != nil {
		return nil, status.Errorf(500, "failed to get devices: %v", err)
	}

	var protoDevices []*pb.Device
	for _, d := range devices {
		protoDevices = append(protoDevices, &pb.Device{
			DeviceType: d.DeviceType,
			DevicePath: d.DevicePath,
			VendorId:   d.VendorID,
			ProductId:  d.ProductID,
		})
	}

	return &pb.DeviceList{Devices: protoDevices}, nil
}
