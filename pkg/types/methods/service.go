package methods

type ServiceCategory struct {
	GreenAPI GreenAPIInterface
}

func (c ServiceCategory) CheckWhatsapp(phoneNumber int) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "CheckWhatsapp", map[string]interface{}{
		"phoneNumber": phoneNumber,
	}, "")
}

func (c ServiceCategory) GetAvatar(chatId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "GetAvatar", map[string]interface{}{
		"chatId": chatId,
	}, "")
}

func (c ServiceCategory) GetContacts() (map[string]interface{}, error) {
	return c.GreenAPI.Request("GET", "GetContacts", nil, "")
}

func (c ServiceCategory) GetContactInfo(chatId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "getContactInfo", map[string]interface{}{
		"chatId": chatId,
	}, "")
}

func (c ServiceCategory) DeleteMessage(chatId, idMessage string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "deleteMessage", map[string]interface{}{
		"chatId":    chatId,
		"idMessage": idMessage,
	}, "")
}

func (c ServiceCategory) ArchiveChat(chatId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "archiveChat", map[string]interface{}{
		"chatId": chatId,
	}, "")
}

func (c ServiceCategory) UnarchiveChat(chatId string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "unarchiveChat", map[string]interface{}{
		"chatId": chatId,
	}, "")
}

func (c ServiceCategory) SetDisappearingChat(parameters map[string]interface{}) (map[string]interface{}, error) {
	method := "GET"
	if parameters != nil {
		method = "POST"
	}

	return c.GreenAPI.Request(method, "setDisappearingChat", parameters, "")
}
