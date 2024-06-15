package http

import (
	"errors"
	"github.com/elephant315/devtree-protobuf/internal/model"
	pb "github.com/elephant315/devtree-protobuf/proto"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockDeviceService struct {
	devices []*model.Device
	err     error
}

func (m *MockDeviceService) GetDevices() ([]*model.Device, error) {
	return m.devices, m.err
}

func TestGetDevices(t *testing.T) {
	// Prepare mock data
	devices := []*model.Device{
		{DeviceType: "mouse", DevicePath: "/dev/input/mouse0", VendorID: "1234", ProductID: "5678"},
		{DeviceType: "keyboard", DevicePath: "/dev/input/keyboard0", VendorID: "1234", ProductID: "5678"},
	}
	mockService := &MockDeviceService{devices: devices}

	e := echo.New()
	handler := NewDeviceHandler(mockService)
	req := httptest.NewRequest(http.MethodGet, "/devices", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the handler
	if assert.NoError(t, handler.GetDevices(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "application/x-protobuf", rec.Header().Get("Content-Type"))

		// Parse the response
		var deviceList pb.DeviceList
		err := proto.Unmarshal(rec.Body.Bytes(), &deviceList)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(deviceList.Devices))
		assert.Equal(t, devices[0].DeviceType, deviceList.Devices[0].DeviceType)
		assert.Equal(t, devices[0].DevicePath, deviceList.Devices[0].DevicePath)
		assert.Equal(t, devices[0].VendorID, deviceList.Devices[0].VendorId)
		assert.Equal(t, devices[0].ProductID, deviceList.Devices[0].ProductId)
	}
}

func TestGetDevices_Error(t *testing.T) {
	mockService := &MockDeviceService{err: errors.New("error")}

	e := echo.New()
	handler := NewDeviceHandler(mockService)
	req := httptest.NewRequest(http.MethodGet, "/devices", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the handler
	if assert.NoError(t, handler.GetDevices(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, "application/x-protobuf", rec.Header().Get("Content-Type"))

		// Parse the error response
		var errorMessage pb.Error
		err := proto.Unmarshal(rec.Body.Bytes(), &errorMessage)
		assert.NoError(t, err)
		assert.Equal(t, "error", errorMessage.Message)
	}
}
