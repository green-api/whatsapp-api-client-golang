package methods

type AccountCategory struct {
	GreenAPI GreenAPIInterface
}

func (c AccountCategory) GetSettings() map[string]interface{} {
	return c.GreenAPI.Request("GET", "GetSettings", nil, "")
}

func (c AccountCategory) SetSettings(parameters map[string]interface{}) map[string]interface{} {
	method := "GET"
	if parameters != nil {
		method = "POST"
	}

	return c.GreenAPI.Request(method, "SetSettings", parameters, "")
}

func (c AccountCategory) SetSystemProxy() map[string]interface{} {
	return c.GreenAPI.Request("GET", "SetSystemProxy", nil, "")
}

func (c AccountCategory) GetStateInstance() map[string]interface{} {
	return c.GreenAPI.Request("GET", "getStateInstance", nil, "")
}

func (c AccountCategory) GetStatusInstance() map[string]interface{} {
	return c.GreenAPI.Request("GET", "getStatusInstance", nil, "")
}

func (c AccountCategory) Reboot() map[string]interface{} {
	return c.GreenAPI.Request("GET", "Reboot", nil, "")
}

func (c AccountCategory) Logout() map[string]interface{} {
	return c.GreenAPI.Request("GET", "Logout", nil, "")
}

func (c AccountCategory) QR() map[string]interface{} {
	return c.GreenAPI.Request("GET", "qr", nil, "")
}

func (c AccountCategory) SetProfilePicture(filePath string) map[string]interface{} {
	return c.GreenAPI.Request("POST", "setProfilePicture", nil, filePath)
}
