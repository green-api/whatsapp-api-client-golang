package methods

type DeviceCategory struct {
	GreenAPI GreenAPIInterface
}

// GetDeviceInfo is designed to get information about the device (phone) on which the WhatsApp Business application is running
func (c DeviceCategory) GetDeviceInfo() (map[string]interface{}, error) {
	return c.GreenAPI.Request("GET", "GetDeviceInfo", nil, "")
}
