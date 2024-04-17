package methods

import (
	"strconv"
)

type StatusCategory struct {
	GreenAPI GreenAPIInterface
}

// SentTextStatus is aimed for sending a text status
// https://green-api.com/en/docs/api/statuses/SendTextStatus/
func (c StatusCategory) SendTextStatus(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "sendTextStatus", parameters, "")
}

// SendVoiceStatus is aimed for sending a voice status
// https://green-api.com/en/docs/api/statuses/SendVoiceStatus/
func (c StatusCategory) SendVoiceStatus(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "sendVoiceStatus", parameters, "")
}

// SendMediaStatus is aimed for sending a pictures or video status
// https://green-api.com/en/docs/api/statuses/SendMediaStatus/
func (c StatusCategory) SendMediaStatus(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "sendMediaStatus", parameters, "")
}

// GetOutgoingStatuses returns the outgoing statuses of the account
// https://green-api.com/en/docs/api/statuses/GetOutgoingStatuses/
func (c StatusCategory) GetOutgoingStatuses(minutes int) ([]interface{}, error) {
	data := map[string]interface{}{"minutes": strconv.Itoa(minutes)}
	if minutes == 0 {
		data = map[string]interface{}{}
	}
	return c.GreenAPI.ArrayRequest("GET", "getOutgoingStatuses", data, "")
}

// GetIncomingStatuses returns the incoming status messages of the account
// https://green-api.com/en/docs/api/statuses/GetIncomingStatuses/
func (c StatusCategory) GetIncomingStatuses(minutes int) ([]interface{}, error) {
	data := map[string]interface{}{"minutes": strconv.Itoa(minutes)}
	if minutes == 0 {
		data = map[string]interface{}{}
	}
	return c.GreenAPI.ArrayRequest("GET", "getIncomingStatuses", data, "")
}

// GetStatusStatistic returns an array of recipients marked for a given status.
// https://green-api.com/en/docs/api/statuses/GetStatusStatistic/
func (c StatusCategory) GetStatusStatistic(idMessage string) ([]interface{}, error) {
	return c.GreenAPI.ArrayRequest("GET", "getStatusStatistic", map[string]interface{}{
		"idMessage": idMessage,
	}, "")
}

// DeleteStatus is aimed for deleting status.
// https://green-api.com/en/docs/api/statuses/DeleteStatus/
func (c StatusCategory) DeleteStatus(idMessage string) (map[string]interface{}, error) {
	_, err := c.GreenAPI.Request("POST", "deleteStatus", map[string]interface{}{
		"idMessage": idMessage,
	}, "")
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{"deletedStatus": idMessage}, nil
}
