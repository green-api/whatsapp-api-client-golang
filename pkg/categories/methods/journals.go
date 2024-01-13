package methods

type JournalsCategory struct {
	GreenAPI GreenAPIInterface
}

// GetChatHistory returns the chat message history.
func (c JournalsCategory) GetChatHistory(parameters map[string]interface{}) ([]map[string]interface{}, error) {
	return c.GreenAPI.ArrayRequest("POST", "getChatHistory", parameters, "")
}

// GetMessage returns a chat message.
func (c JournalsCategory) GetMessage(chatId, idMessage string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "getMessage", map[string]interface{}{
		"chatId":    chatId,
		"idMessage": idMessage,
	}, "")
}

// LastIncomingMessages returns the most recent incoming messages
// of the account.
func (c JournalsCategory) LastIncomingMessages(parameters map[string]interface{}) ([]map[string]interface{}, error) {
	return c.GreenAPI.ArrayRequest("GET", "lastIncomingMessages", parameters, "")
}

// LastOutgoingMessages returns the last sent messages of the account.
func (c JournalsCategory) LastOutgoingMessages(parameters map[string]interface{}) ([]map[string]interface{}, error) {
	return c.GreenAPI.ArrayRequest("GET", "lastOutgoingMessages", parameters, "")
}
