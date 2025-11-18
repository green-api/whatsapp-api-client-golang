package methods

type JournalsCategory struct {
	GreenAPI GreenAPIInterface
}

// GetChatHistory returns the chat message history.
// https://green-api.com/en/docs/api/journals/GetChatHistory/
func (c JournalsCategory) GetChatHistory(parameters map[string]any) ([]any, error) {
	return c.GreenAPI.ArrayRequest("POST", "getChatHistory", parameters, "")
}

// GetMessage returns a chat message.
// https://green-api.com/en/docs/api/journals/GetMessage/
func (c JournalsCategory) GetMessage(chatId, idMessage string) (map[string]any, error) {
	return c.GreenAPI.Request("POST", "getMessage", map[string]any{
		"chatId":    chatId,
		"idMessage": idMessage,
	}, "")
}

// LastIncomingMessages returns the most recent incoming messages
// of the account.
// https://green-api.com/en/docs/api/journals/LastIncomingMessages/
func (c JournalsCategory) LastIncomingMessages(parameters map[string]any) ([]any, error) {
	return c.GreenAPI.ArrayRequest("GET", "lastIncomingMessages", parameters, "")
}

// LastOutgoingMessages returns the last sent messages of the account.
// https://green-api.com/en/docs/api/journals/LastOutgoingMessages/
func (c JournalsCategory) LastOutgoingMessages(parameters map[string]any) ([]any, error) {
	return c.GreenAPI.ArrayRequest("GET", "lastOutgoingMessages", parameters, "")
}
