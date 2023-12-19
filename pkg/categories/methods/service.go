package methods

type ServiceCategory struct {
	GreenAPI GreenAPIInterface
}

// CheckWhatsapp checks if there is a WhatsApp account on the phone number.
func (c ServiceCategory) CheckWhatsapp(phoneNumber int) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "checkWhatsapp", map[string]interface{}{
		"phoneNumber": phoneNumber,
	}, "")
}

// GetAvatar returns the avatar of the correspondent or group chat.
func (c ServiceCategory) GetAvatar(chatId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "getAvatar", map[string]interface{}{
		"chatId": chatId,
	}, "")
}

// GetContacts is designed to get a list of contacts of the current account.
func (c ServiceCategory) GetContacts() ([]interface{}, error) {
	return c.GreenAPI.ArrayRequest("GET", "getContacts", nil, "")
}

// GetContactInfo is designed to obtain information about the contact.
func (c ServiceCategory) GetContactInfo(chatId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "getContactInfo", map[string]interface{}{
		"chatId": chatId,
	}, "")
}

// DeleteMessage deletes the message from chat.
func (c ServiceCategory) DeleteMessage(chatId, idMessage string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "deleteMessage", map[string]interface{}{
		"chatId":    chatId,
		"idMessage": idMessage,
	}, "")
}

// ArchiveChat archives the chat.
func (c ServiceCategory) ArchiveChat(chatId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "archiveChat", map[string]interface{}{
		"chatId": chatId,
	}, "")
}

// UnarchiveChat unarchives the chat.
func (c ServiceCategory) UnarchiveChat(chatId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "unarchiveChat", map[string]interface{}{
		"chatId": chatId,
	}, "")
}

// SetDisappearingChat is designed to change the settings
// of disappearing messages in chats.
func (c ServiceCategory) SetDisappearingChat(parameters map[string]interface{}) (map[string]interface{}, error) {
	method := "GET"
	if parameters != nil {
		method = "POST"
	}

	return c.GreenAPI.Request(method, "setDisappearingChat", parameters, "")
}
