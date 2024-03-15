package methods

type SendingCategory struct {
	GreenAPI GreenAPIInterface
}

// SendMessage is designed to send a text message to a personal or group chat.
// https://green-api.com/en/docs/api/sending/SendMessage/
func (c SendingCategory) SendMessage(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "sendMessage", parameters, "")
}

// SendButtons is designed to send a message with buttons
// to a personal or group chat.
// https://green-api.com/en/docs/api/sending/SendButtons/
func (c SendingCategory) SendButtons(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "sendButtons", parameters, "")
}

// SendTemplateButtons is designed to send a message with interactive buttons
// from the list of templates in a personal or group chat.
// https://green-api.com/en/docs/api/sending/SendTemplateButtons/
func (c SendingCategory) SendTemplateButtons(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "sendTemplateButtons", parameters, "")
}

// SendListMessage is designed to send a message with a selection button
// from a list of values to a personal or group chat.
// https://green-api.com/en/docs/api/sending/SendListMessage/
func (c SendingCategory) SendListMessage(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "sendListMessage", parameters, "")
}

// SendFileByUpload is designed to send a file loaded through a form (form-data).
// https://green-api.com/en/docs/api/sending/SendFileByUpload/
func (c SendingCategory) SendFileByUpload(filePath string, parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "sendFileByUpload", parameters, filePath)
}

// SendFileByUrl is designed to send a file downloaded via a link.
// https://green-api.com/en/docs/api/sending/SendFileByUrl/
func (c SendingCategory) SendFileByUrl(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "sendFileByUrl", parameters, "")
}

// SendLocation is designed to send a geolocation message.
// https://green-api.com/en/docs/api/sending/SendLocation/
func (c SendingCategory) SendLocation(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "sendLocation", parameters, "")
}

// SendContact is for sending a message with a contact.
// https://green-api.com/en/docs/api/sending/SendContact/
func (c SendingCategory) SendContact(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "sendContact", parameters, "")
}

// The method is deprecated. Please use sendMessage.
func (c SendingCategory) SendLink(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "sendLink", parameters, "")
}

// ForwardMessages is designed for forwarding messages
// to a personal or group chat.
// https://green-api.com/en/docs/api/sending/ForwardMessages/
func (c SendingCategory) ForwardMessages(chatId, chatIdFrom string, messages []string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "forwardMessages", map[string]interface{}{
		"chatId":     chatId,
		"chatIdFrom": chatIdFrom,
		"messages":   messages,
	}, "")
}

// UploadFile allows you to upload a file from the local file system,
// which can later be sent using the SendFileByUrl method.
// https://green-api.com/en/docs/api/sending/UploadFile/
func (c SendingCategory) UploadFile(filePath string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "uploadFile", nil, filePath)
}

// SendPoll is designed for sending messages with a poll
// to a private or group chat.
// https://green-api.com/en/docs/api/sending/SendPoll/
func (c SendingCategory) SendPoll(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "sendPoll", parameters, "")
}
