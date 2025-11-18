package methods

type AccountCategory struct {
	GreenAPI GreenAPIInterface
}

// GetSettings is designed to get the current settings of the account.
// https://green-api.com/en/docs/api/account/GetSettings/
func (c AccountCategory) GetSettings() (map[string]any, error) {
	return c.GreenAPI.Request("GET", "getSettings", nil, "")
}

// GetWaSettings is designed to get information about the WhatsApp account.
// https://green-api.com/en/docs/api/account/GetWaSettings/
func (c AccountCategory) GetWaSettings() (map[string]any, error) {
	return c.GreenAPI.Request("GET", "getWaSettings", nil, "")
}

// SetSettings is for setting the account settings.
// https://green-api.com/en/docs/api/account/SetSettings/
func (c AccountCategory) SetSettings(parameters map[string]any) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "setSettings", parameters, "")
}

// GetStateInstance is designed to get the state of the account.
// https://green-api.com/en/docs/api/account/GetStateInstance/
func (c AccountCategory) GetStateInstance() (map[string]any, error) {
	return c.GreenAPI.Request("GET", "getStateInstance", nil, "")
}

// GetStatusInstance is designed to get the socket connection state
// of the account instance with WhatsApp.
// https://green-api.com/en/docs/api/account/GetStatusInstance/
func (c AccountCategory) GetStatusInstance() (map[string]any, error) {
	return c.GreenAPI.Request("GET", "getStatusInstance", nil, "")
}

// Reboot is designed to restart the account.
// https://green-api.com/en/docs/api/account/Reboot/
func (c AccountCategory) Reboot() (map[string]any, error) {
	return c.GreenAPI.Request("GET", "reboot", nil, "")
}

// Logout is designed to unlogin the account.
// https://green-api.com/en/docs/api/account/Logout/
func (c AccountCategory) Logout() (map[string]any, error) {
	return c.GreenAPI.Request("GET", "logout", nil, "")
}

// QR is designed to get a QR code.
// https://green-api.com/en/docs/api/account/QR/
func (c AccountCategory) QR() (map[string]any, error) {
	return c.GreenAPI.Request("GET", "qr", nil, "")
}

// SetProfilePicture is designed to set the avatar of the account.
// https://green-api.com/en/docs/api/account/SetProfilePicture/
func (c AccountCategory) SetProfilePicture(filePath string) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "setProfilePicture", nil, filePath)
}

// GetAuthorizationCode is designed to authorize an instance by phone number.
// https://green-api.com/en/docs/api/account/GetAuthorizationCode/
func (c AccountCategory) GetAuthorizationCode(phoneNumber int) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "getAuthorizationCode", map[string]any{
		"phoneNumber": phoneNumber,
	}, "")
}
