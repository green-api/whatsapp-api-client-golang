package methods

type DeviceCategory struct {
	GreenAPI GreenAPIInterface
}

func (c DeviceCategory) GetDeviceInfo() map[string]interface{} {
	return c.GreenAPI.Request("GET", "GetDeviceInfo", nil, "")
}
