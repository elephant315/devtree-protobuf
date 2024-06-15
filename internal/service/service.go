package service

import (
	"github.com/elephant315/devtree-protobuf/internal/model"
	"github.com/elephant315/devtree-protobuf/internal/ports"
)

type DeviceService struct {
	adapter ports.DeviceAdapter
}

func NewDeviceService(adapter ports.DeviceAdapter) ports.DeviceService {
	return &DeviceService{
		adapter: adapter,
	}
}

func (s *DeviceService) GetDevices() ([]*model.Device, error) {
	return s.adapter.GetDevices()
}