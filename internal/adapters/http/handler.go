package http

import (
	"github.com/elephant315/devtree-protobuf/internal/ports"
	pb "github.com/elephant315/devtree-protobuf/proto"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/proto"
)

type DeviceHandler struct {
	service ports.DeviceService
}

func NewDeviceHandler(service ports.DeviceService) *DeviceHandler {
	return &DeviceHandler{
		service: service,
	}
}

func (h *DeviceHandler) GetDevices(c echo.Context) error {
	devices, err := h.service.GetDevices()
	if err != nil {
		return h.respondWithError(c, err)
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

	deviceList := &pb.DeviceList{Devices: protoDevices}
	data, err := proto.Marshal(deviceList)
	if err != nil {
		return h.respondWithError(c, err)
	}
	return c.Blob(http.StatusOK, "application/x-protobuf", data)
}

func (h *DeviceHandler) respondWithError(c echo.Context, err error) error {
	errorMessage := &pb.Error{Message: err.Error()}
	data, marshalErr := proto.Marshal(errorMessage)
	if marshalErr != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.Blob(http.StatusInternalServerError, "application/x-protobuf", data)
}
