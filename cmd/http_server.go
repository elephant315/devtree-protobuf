package main

import (
	"github.com/elephant315/devtree-protobuf/internal/adapters/device"
	"github.com/elephant315/devtree-protobuf/internal/adapters/http"
	"github.com/elephant315/devtree-protobuf/internal/service"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	deviceAdapter := device.NewDeviceAdapter()
	deviceService := service.NewDeviceService(deviceAdapter)
	handler := http.NewDeviceHandler(deviceService)

	e.GET("/devices", handler.GetDevices)

	e.Logger.Fatal(e.Start(":8080"))
}
