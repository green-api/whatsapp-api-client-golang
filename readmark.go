package greenapi

import "encoding/json"

type ReadMarkCategory struct {
	GreenAPI GreenAPIInterface
}

// ------------------------------------------------------------------ ReadChat block

type RequestReadChat struct {
	ChatId    string `json:"chatId"`
	IdMessage string `json:"idMessage,omitempty"`
}

type ReadChatOption func(*RequestReadChat) error

// ID of the incoming message to be marked as read. If not specified, then all unread messages in the chat will be marked as read.
func OptionalIdMessage(idMessage string) ReadChatOption {
	return func(r *RequestReadChat) error {
		r.IdMessage = idMessage
		return nil
	}
}

// Marking messages in a chat as read.
//
// https://green-api.com/en/docs/api/marks/ReadChat/
//
// Add optional arguments by passing these functions:
//  OptionalIdMessage(idMessage string) <- ID of the incoming message to be marked as read. If not specified, then all unread messages in the chat will be marked as read.
func (c ReadMarkCategory) ReadChat(chatId string, options ...ReadChatOption) (*APIResponse, error) {
	err := ValidateChatId(chatId)
	if err!=nil {
		return nil, err
	}
	
	r := &RequestReadChat{
		ChatId: chatId,
	}

	for _, o := range options {
		o(r)
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return c.GreenAPI.Request("POST", "readChat", jsonData)
}
