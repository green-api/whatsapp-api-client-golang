package methods

type ServiceCategory struct {
	GreenAPI GreenAPIInterface
}

// CheckWhatsapp checks if there is a WhatsApp account on the phone number.
// https://green-api.com/en/docs/api/service/CheckWhatsapp/
func (c ServiceCategory) CheckWhatsapp(phoneNumber int) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "checkWhatsapp", map[string]any{
		"phoneNumber": phoneNumber,
	}, "")
}

// GetAvatar returns the avatar of the correspondent or group chat.
// https://green-api.com/en/docs/api/service/GetAvatar/
func (c ServiceCategory) GetAvatar(chatId string) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "getAvatar", map[string]any{
		"chatId": chatId,
	}, "")
}

// GetContacts is designed to get a list of contacts of the current account.
// https://green-api.com/en/docs/api/service/GetContacts/
func (c ServiceCategory) GetContacts() ([]any, error) {
	return c.GreenAPI.ArrayRequest("GET", "getContacts", nil, "")
}

// GetContactInfo is designed to obtain information about the contact.
// https://green-api.com/en/docs/api/service/GetContactInfo/
func (c ServiceCategory) GetContactInfo(chatId string) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "getContactInfo", map[string]any{
		"chatId": chatId,
	}, "")
}

// DeleteMessage deletes the message from chat.
// https://green-api.com/en/docs/api/service/deleteMessage/
func (c ServiceCategory) DeleteMessage(chatId, idMessage string) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "deleteMessage", map[string]any{
		"chatId":    chatId,
		"idMessage": idMessage,
	}, "")
}

// ArchiveChat archives the chat.
// https://green-api.com/en/docs/api/service/archiveChat/
func (c ServiceCategory) ArchiveChat(chatId string) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "archiveChat", map[string]any{
		"chatId": chatId,
	}, "")
}

// UnarchiveChat unarchives the chat.
// https://green-api.com/en/docs/api/service/unarchiveChat/
func (c ServiceCategory) UnarchiveChat(chatId string) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "unarchiveChat", map[string]any{
		"chatId": chatId,
	}, "")
}

// SetDisappearingChat is designed to change the settings of disappearing messages in chats.
// https://green-api.com/en/docs/api/service/SetDisappearingChat/
func (c ServiceCategory) SetDisappearingChat(parameters map[string]any) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "setDisappearingChat", parameters, "")
}

// SendTyping is designed to send a notification about typing or recording audio in a chat.
// https://green-api.com/en/docs/api/service/SendTyping/
func (c ServiceCategory) SendTyping(chatId string) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "sendTyping", map[string]any{
		"chatId": chatId,
	}, "")
}
