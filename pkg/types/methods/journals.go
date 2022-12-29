package methods

type JournalsCategory struct {
	GreenAPI GreenAPIInterface
}

func (c JournalsCategory) GetChatHistory(parameters map[string]interface{}) map[string]interface{} {
	return c.GreenAPI.Request("POST", "GetChatHistory", parameters, "")
}

func (c JournalsCategory) GetMessage(chatId, idMessage string) map[string]interface{} {
	return c.GreenAPI.Request("POST", "getMessage", map[string]interface{}{
		"chatId":    chatId,
		"idMessage": idMessage,
	}, "")
}

func (c JournalsCategory) LastIncomingMessages(parameters map[string]interface{}) map[string]interface{} {
	return c.GreenAPI.Request("GET", "lastIncomingMessages", parameters, "")
}

func (c JournalsCategory) LastOutgoingMessages(parameters map[string]interface{}) map[string]interface{} {
	return c.GreenAPI.Request("GET", "LastOutgoingMessages", parameters, "")
}
