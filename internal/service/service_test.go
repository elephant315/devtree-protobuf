package service

import (
	"github.com/elephant315/devtree-protobuf/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDeviceAdapter is a mock implementation of the DeviceAdapter interface
type MockDeviceAdapter struct {
	mock.Mock
}

func (m *MockDeviceAdapter) GetDevices() ([]*model.Device, error) {
	args := m.Called()
	return args.Get(0).([]*model.Device), args.Error(1)
}

func TestDeviceService_GetDevices(t *testing.T) {
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

	service := NewDeviceService(mockAdapter)

	devices, err := service.GetDevices()
	assert.NoError(t, err)
	assert.Equal(t, expectedDevices, devices)

	// Assert that the expectations were met
	mockAdapter.AssertExpectations(t)
}
