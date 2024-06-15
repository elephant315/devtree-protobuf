package device

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDevices(t *testing.T) {
	adapter := NewDeviceAdapter()

	devices, err := adapter.GetDevices()

	assert.NoError(t, err)
	assert.NotNil(t, devices)

	if len(devices) > 0 {
		assert.NotEmpty(t, devices[0].DeviceType)
		assert.NotEmpty(t, devices[0].DevicePath)
		assert.NotEmpty(t, devices[0].VendorID)
		assert.NotEmpty(t, devices[0].ProductID)
	}
}
