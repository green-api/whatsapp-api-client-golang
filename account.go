package greenapi

import (
	"encoding/json"
	"fmt"
)

type AccountCategory struct {
	GreenAPI GreenAPIInterface
}

func (c AccountCategory) GetSettings() (interface{}, error) {
	return c.GreenAPI.Request("GET", "getSettings", nil)
}

// func (c AccountCategory) GetInstances() (interface{}, error) {
// 	return c.GreenAPI.Request("GET", "getInstances", map[string]interface{}{}) //обязательно передавать пустой интерфейс вместо nil
// }

// OptionsSetSettings contains available parameters for SetSettings
type requestSetSettings struct {
	WebhookUrl                        *string `json:"webhookUrl,omitempty"`
	WebhookUrlToken                   *string `json:"webhookUrlToken,omitempty"`
	DelaySendMessagesMilliseconds     *int    `json:"delaySendMessagesMilliseconds,omitempty"`
	MarkIncomingMessagesReaded        string  `json:"markIncomingMessagesReaded,omitempty"`
	MarkIncomingMessagesReadedOnReply string  `json:"markIncomingMessagesReadedOnReply,omitempty"`
	OutgoingWebhook                   string  `json:"outgoingWebhook,omitempty"`
	OutgoingMessageWebhook            string  `json:"outgoingMessageWebhook,omitempty"`
	OutgoingAPIMessageWebhook         string  `json:"outgoingAPIMessageWebhook,omitempty"`
	StateWebhook                      string  `json:"stateWebhook,omitempty"`
	IncomingWebhook                   string  `json:"incomingWebhook,omitempty"`
	DeviceWebhook                     string  `json:"deviceWebhook,omitempty"`
	KeepOnlineStatus                  string  `json:"keepOnlineStatus,omitempty"`
	PollMessageWebhook                string  `json:"pollMessageWebhook,omitempty"`
	IncomingBlockWebhook              string  `json:"incomingBlockWebhook,omitempty"`
	IncomingCallWebhook               string  `json:"incomingCallWebhook,omitempty"`
}

type AccountOption func(*requestSetSettings)

func WithWebhookUrl(webhookUrl string) AccountOption {
	return func(r *requestSetSettings) {
		r.WebhookUrl = &webhookUrl
	}
}

func WithWebhookUrlToken(webhookUrlToken string) AccountOption {
	return func(r *requestSetSettings) {
		r.WebhookUrlToken = &webhookUrlToken
	}
}

func WithDelaySendMesssages(delaySendMessagesMilliseconds int) AccountOption {
	return func(r *requestSetSettings) {
		r.DelaySendMessagesMilliseconds = &delaySendMessagesMilliseconds
	}
}

func WithMarkIncomingMessagesRead(markIncomingMessagesReaded string) AccountOption {
	return func(r *requestSetSettings) {
		r.MarkIncomingMessagesReaded = markIncomingMessagesReaded
	}
}

func WithMarkIncomingMessagesReadOnReply(markIncomingMessagesReadedOnReply string) AccountOption {
	return func(r *requestSetSettings) {
		r.MarkIncomingMessagesReadedOnReply = markIncomingMessagesReadedOnReply
	}
}

func WithOutgoingWebhook(outgoingWebhook string) AccountOption {
	return func(r *requestSetSettings) {
		r.OutgoingWebhook = outgoingWebhook
	}
}

func WithOutgoingMessageWebhook(outgoingMessageWebhook string) AccountOption {
	return func(r *requestSetSettings) {
		r.OutgoingMessageWebhook = outgoingMessageWebhook
	}
}

func WithOutgoingAPIMessageWebhook(outgoingAPIMessageWebhook string) AccountOption {
	return func(r *requestSetSettings) {
		r.OutgoingAPIMessageWebhook = outgoingAPIMessageWebhook
	}
}

func WithStateWebhook(stateWebhook string) AccountOption {
	return func(r *requestSetSettings) {
		r.StateWebhook = stateWebhook
	}
}

func WithIncomingWebhook(incomingWebhook string) AccountOption {
	return func(r *requestSetSettings) {
		r.IncomingWebhook = incomingWebhook
	}
}

func WithDeviceWebhook(deviceWebhook string) AccountOption {
	return func(r *requestSetSettings) {
		r.DeviceWebhook = deviceWebhook
	}
}

func WithKeepOnlineStatus(keepOnlineStatus string) AccountOption {
	return func(r *requestSetSettings) {
		r.KeepOnlineStatus = keepOnlineStatus
	}
}

func WithPollMessageWebhook(pollMessageWebhook string) AccountOption {
	return func(r *requestSetSettings) {
		r.PollMessageWebhook = pollMessageWebhook
	}
}

func WithIncomingBlockWebhook(incomingBlockWebhook string) AccountOption {
	return func(r *requestSetSettings) {
		r.IncomingBlockWebhook = incomingBlockWebhook
	}
}

func WithIncomingCallWebhook(incomingCallWebhook string) AccountOption {
	return func(r *requestSetSettings) {
		r.IncomingCallWebhook = incomingCallWebhook
	}
}

func (c AccountCategory) SetSettings(options ...AccountOption) (interface{}, error) {

	r := &requestSetSettings{}
	for _, o := range options {
		o(r)
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	var payload map[string]interface{}

	if err := json.Unmarshal(jsonData, &payload); err != nil {
		return nil, err
	}

	fmt.Println(payload)
	return c.GreenAPI.Request("POST", "setSettings", payload)
}