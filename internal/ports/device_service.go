package ports

import "github.com/elephant315/devtree-protobuf/internal/model"

type DeviceAdapter interface {
	GetDevices() ([]*model.Device, error)
}

type DeviceService interface {
	GetDevices() ([]*model.Device, error)
}