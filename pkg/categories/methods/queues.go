package methods

type QueuesCategory struct {
	GreenAPI GreenAPIInterface
}

// ShowMessagesQueue is designed to get the list of messages
// that are in the queue to be sent.
// https://green-api.com/en/docs/api/queues/ShowMessagesQueue/
func (c QueuesCategory) ShowMessagesQueue() ([]interface{}, error) {
	return c.GreenAPI.ArrayRequest("GET", "showMessagesQueue", nil, "")
}

// ClearMessagesQueue is designed to clear the queue of messages to be sent.
// https://green-api.com/en/docs/api/queues/ClearMessagesQueue/
func (c QueuesCategory) ClearMessagesQueue() (map[string]interface{}, error) {
	return c.GreenAPI.Request("GET", "clearMessagesQueue", nil, "")
}
