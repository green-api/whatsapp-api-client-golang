package methods

type QueuesCategory struct {
	GreenAPI GreenAPIInterface
}

func (c QueuesCategory) ShowMessagesQueue() map[string]interface{} {
	return c.GreenAPI.Request("GET", "ShowMessagesQueue", nil, "")
}

func (c QueuesCategory) ClearMessagesQueue() map[string]interface{} {
	return c.GreenAPI.Request("GET", "ClearMessagesQueue", nil, "")
}
