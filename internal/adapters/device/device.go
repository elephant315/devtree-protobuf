package device

import (
	"fmt"
	"github.com/citilinkru/libudev"
	"github.com/citilinkru/libudev/matcher"
	"github.com/elephant315/devtree-protobuf/internal/model"
)

type DeviceAdapter struct{
	// here could be some device filtering params!
}

func NewDeviceAdapter() *DeviceAdapter {
	return &DeviceAdapter{}
}

func (da *DeviceAdapter) GetDevices() ([]*model.Device, error) {
	s := libudev.NewScanner()
	err, devices := s.ScanDevices()

	if err != nil {
		return nil, err
	}
	m := matcher.NewMatcher()
	m.SetStrategy(matcher.StrategyOr)
	// Edit this for proper device type selection
	busTypes := []string{"usb", "pci", "pci_express", "bluetooth", "serio", "hid", "platform"}
	for _, bus := range busTypes {
		m.AddRule(matcher.NewRuleEnv("ID_BUS", bus))
	}
	// Additionally, include all inputs and audio
	m.AddRule(matcher.NewRuleEnv("ID_INPUT", "1"))
	m.AddRule(matcher.NewRuleEnv("ID_AUDIO", "1"))

	filteredDevices := m.Match(devices)
	var result []*model.Device
	for _, d := range filteredDevices {
		dev := &model.Device{
			DeviceType:     fmt.Sprintf("%s %s",
							defaultIfEmpty(d.Env["DEVNAME"], "[noname]"),
							defaultIfEmpty(d.Env["ID_TYPE"], "[unknown]")),
			DevicePath:     defaultIfEmpty(d.Devpath, "[no_path]"),
			VendorID:       defaultIfEmpty(d.Env["ID_VENDOR_ID"], "[empty]"),
			ProductID:      defaultIfEmpty(d.Env["ID_MODEL_ID"], "[empty]"),
		}
		result = append(result, dev)
	}
	return result, nil
}

func defaultIfEmpty(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}