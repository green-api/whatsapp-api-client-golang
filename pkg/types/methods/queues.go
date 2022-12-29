package methods

type QueuesCategory struct {
	GreenAPI GreenAPIInterface
}

// ShowMessagesQueue is designed to get the list of messages that are in the queue to be sent
func (c QueuesCategory) ShowMessagesQueue() (map[string]interface{}, error) {
	return c.GreenAPI.Request("GET", "ShowMessagesQueue", nil, "")
}

// ClearMessagesQueue is designed to clear the queue of messages to be sent
func (c QueuesCategory) ClearMessagesQueue() (map[string]interface{}, error) {
	return c.GreenAPI.Request("GET", "ClearMessagesQueue", nil, "")
}
