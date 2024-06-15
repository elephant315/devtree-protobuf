package grpc

import (
	"context"
	"net"
	"testing"
	"github.com/elephant315/devtree-protobuf/internal/model"
	"github.com/elephant315/devtree-protobuf/internal/service"
	"github.com/elephant315/devtree-protobuf/proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

// MockDeviceAdapter is a mock implementation of the DeviceAdapter interface
type MockDeviceAdapter struct {
	mock.Mock
}

func (m *MockDeviceAdapter) GetDevices() ([]*model.Device, error) {
	args := m.Called()
	return args.Get(0).([]*model.Device), args.Error(1)
}

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGRPCServer_GetDevices(t *testing.T) {
	mockAdapter := new(MockDeviceAdapter)

	// Define the expected output
	expectedDevices := []*model.Device{
		{
			DeviceType: "mouse",
			DevicePath: "/dev/input/mouse0",
			VendorID:   "1234",
			ProductID:  "5678",
		},
		{
			DeviceType: "keyboard",
			DevicePath: "/dev/input/keyboard0",
			VendorID:   "8765",
			ProductID:  "4321",
		},
	}

	// Setup the expected calls and return values
	mockAdapter.On("GetDevices").Return(expectedDevices, nil)

	// Create the device service with the mock adapter
	deviceService := service.NewDeviceService(mockAdapter)

	// Create the gRPC server with the service
	grpcServer := NewGRPCServer(deviceService)

	s := grpc.NewServer()
	proto.RegisterDeviceServiceServer(s, grpcServer)
	go func() {
		if err := s.Serve(lis); err != nil {
			t.Fatalf("Server exited with error: %v", err)
		}
	}()

	// Create a gRPC client and connect to the test server
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := proto.NewDeviceServiceClient(conn)

	// Call the GetDevices method
	req := &proto.GetDevicesRequest{}
	res, err := client.GetDevices(ctx, req)
	assert.NoError(t, err)

	// Verify the response
	assert.Len(t, res.Devices, len(expectedDevices))
	for i, device := range res.Devices {
		assert.Equal(t, expectedDevices[i].DeviceType, device.DeviceType)
		assert.Equal(t, expectedDevices[i].DevicePath, device.DevicePath)
		assert.Equal(t, expectedDevices[i].VendorID, device.VendorId)
		assert.Equal(t, expectedDevices[i].ProductID, device.ProductId)
	}

	// Assert that the expectations were met
	mockAdapter.AssertExpectations(t)
}
