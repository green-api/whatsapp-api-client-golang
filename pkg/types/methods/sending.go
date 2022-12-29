package methods

type SendingCategory struct {
	GreenAPI GreenAPIInterface
}

func (c SendingCategory) SendMessage(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "SendMessage", parameters, "")
}

func (c SendingCategory) SendButtons(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "SendButtons", parameters, "")
}

func (c SendingCategory) SendTemplateButtons(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "sendTemplateButtons", parameters, "")
}

func (c SendingCategory) SendListMessage(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "SendListMessage", parameters, "")
}

func (c SendingCategory) SendFileByUpload(filePath string, parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "SendFileByUpload", parameters, filePath)
}

func (c SendingCategory) SendFileByUrl(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "SendFileByUrl", parameters, "")
}

func (c SendingCategory) SendLocation(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "sendLocation", parameters, "")
}

func (c SendingCategory) SendContact(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "sendContact", parameters, "")
}

func (c SendingCategory) SendLink(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "sendLink", parameters, "")
}
