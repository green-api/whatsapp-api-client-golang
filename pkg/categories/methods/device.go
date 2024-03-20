package methods

type DeviceCategory struct {
	GreenAPI GreenAPIInterface
}

// GetDeviceInfo is designed to get information about the device (phone)
// on which the WhatsApp Business application is running.
// https://green-api.com/en/docs/api/phone/GetDeviceInfo/
func (c DeviceCategory) GetDeviceInfo() (map[string]interface{}, error) {
	return c.GreenAPI.Request("GET", "getDeviceInfo", nil, "")
}
