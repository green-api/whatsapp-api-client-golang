package methods

import "fmt"

type SendingCategory struct {
	GreenAPI GreenAPIInterface
}

// SendMessage is designed to send a text message to a personal or group chat.
// https://green-api.com/en/docs/api/sending/SendMessage/
func (c SendingCategory) SendMessage(parameters map[string]any) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "sendMessage", parameters, "")
}

// The method is deprecated. Please use SendInteractiveButtons.
func (c SendingCategory) SendButtons(parameters map[string]any) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "sendButtons", parameters, "")
}

// The method is deprecated. Please use SendInteractiveButtonsReply.
func (c SendingCategory) SendTemplateButtons(parameters map[string]any) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "sendTemplateButtons", parameters, "")
}

// The method is deprecated. Please use SendMessage.
func (c SendingCategory) SendListMessage(parameters map[string]any) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "sendListMessage", parameters, "")
}

// SendFileByUpload is designed to send a file loaded through a form-data.
// https://green-api.com/en/docs/api/sending/SendFileByUpload/
func (c SendingCategory) SendFileByUpload(filePath string, parameters map[string]any) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "sendFileByUpload", parameters, filePath)
}

// SendFileByUrl is designed to send a file downloaded via a link.
// https://green-api.com/en/docs/api/sending/SendFileByUrl/
func (c SendingCategory) SendFileByUrl(parameters map[string]any) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "sendFileByUrl", parameters, "")
}

// SendLocation is designed to send a geolocation message.
// https://green-api.com/en/docs/api/sending/SendLocation/
func (c SendingCategory) SendLocation(parameters map[string]any) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "sendLocation", parameters, "")
}

// SendContact is for sending a message with a contact.
// https://green-api.com/en/docs/api/sending/SendContact/
func (c SendingCategory) SendContact(parameters map[string]any) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "sendContact", parameters, "")
}

// The method is deprecated. Please use SendMessage.
func (c SendingCategory) SendLink(parameters map[string]any) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "sendLink", parameters, "")
}

// ForwardMessages is designed for forwarding messages
// to a personal or group chat.
// https://green-api.com/en/docs/api/sending/ForwardMessages/
func (c SendingCategory) ForwardMessages(chatId, chatIdFrom string, messages []string) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "forwardMessages", map[string]any{
		"chatId":     chatId,
		"chatIdFrom": chatIdFrom,
		"messages":   messages,
	}, "")
}

// UploadFile allows you to upload a file from the local file system,
// which can later be sent using the SendFileByUrl method.
// https://green-api.com/en/docs/api/sending/UploadFile/
func (c SendingCategory) UploadFile(filePath string) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "uploadFile", nil, filePath)
}

// SendPoll is designed for sending messages with a poll
// to a private or group chat.
// https://green-api.com/en/docs/api/sending/SendPoll/
func (c SendingCategory) SendPoll(parameters map[string]any) (map[string]any, error) {
	message, ok := parameters["message"].(string)
	if !ok {
		return nil, fmt.Errorf("cannot find message paramater")
	}

	if len(message) > 255 {
		return nil, fmt.Errorf("number of characters in message exceeded (more than 255)")
	}

	if parameters == nil {
		return nil, fmt.Errorf("parameters is nil")
	}

	optionsValue, exists := parameters["options"]
	if !exists || optionsValue == nil {
		return nil, fmt.Errorf(`parameters["options"] is nil or not exists`)
	}

	options, ok := optionsValue.([]map[string]any)
	if !ok {
		return nil, fmt.Errorf("options is not of type []map[string]any, got %T", optionsValue)
	}

	if len(options) < 2 {
		return nil, fmt.Errorf("cannot create less than 2 poll options")
	} else if len(options) > 12 {
		return nil, fmt.Errorf("cannot create more than 12 poll options")
	}

	seen := make(map[string]bool)

	for _, option := range options {
		optionValue, ok := option["optionName"].(string)
		if len(optionValue) > 100 {
			return nil, fmt.Errorf("number of characters in optionName exceeded (more than 100)")
		}
		if !ok {
			return nil, fmt.Errorf("option does not have a valid 'optionName'")
		}
		if seen[optionValue] {
			return nil, fmt.Errorf("poll options cannot have duplicates: %s", optionValue)
		}
		seen[optionValue] = true
	}

	return c.GreenAPI.Request("POST", "sendPoll", parameters, "")
}

// SendInteractiveButtons is designed to send a message with interactive buttons
// to a personal or group chat.
// https://green-api.com/en/docs/api/sending/SendInteractiveButtons/
func (c SendingCategory) SendInteractiveButtons(parameters map[string]any) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "sendInteractiveButtons", parameters, "")
}

// SendInteractiveButtonsReply is designed to send a message with interactive reply buttons
// to a personal or group chat.
// https://green-api.com/en/docs/api/sending/SendInteractiveButtons/
func (c SendingCategory) SendInteractiveButtonsReply(parameters map[string]any) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "sendInteractiveButtonsReply", parameters, "")
}
