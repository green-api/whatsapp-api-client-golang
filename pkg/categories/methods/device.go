package methods

import "errors"

type DeviceCategory struct {
	GreenAPI GreenAPIInterface
}

func (c DeviceCategory) GetDeviceInfo() (map[string]any, error) {
	return nil, errors.New("method GetDeviceInfo() is deprecated and disabled")
}
