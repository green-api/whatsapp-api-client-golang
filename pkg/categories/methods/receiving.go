package methods

type ReceivingCategory struct {
	GreenAPI GreenAPIInterface
}

// ReceiveNotification is designed to receive a single incoming notification
// from the notification queue.
// https://green-api.com/en/docs/api/receiving/technology-http-api/ReceiveNotification/
func (c ReceivingCategory) ReceiveNotification() (map[string]interface{}, error) {
	response, err := c.GreenAPI.RawRequest("GET", "receiveNotification", nil, "")

	if response != nil {
		return response.(map[string]interface{}), err
	}

	return nil, err
}

// DeleteNotification is designed to remove an incoming notification
// from the notification queue.
// https://green-api.com/en/docs/api/receiving/technology-http-api/DeleteNotification/
func (c ReceivingCategory) DeleteNotification(receiptId int) (map[string]interface{}, error) {
	return c.GreenAPI.Request("DELETE", "deleteNotification", map[string]interface{}{
		"receiptId": receiptId,
	}, "")
}

// DownloadFile is for downloading received and sent files.
// https://green-api.com/en/docs/api/receiving/files/DownloadFile/
func (c ReceivingCategory) DownloadFile(chatId, idMessage string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "downloadFile", map[string]interface{}{
		"chatId":    chatId,
		"idMessage": idMessage,
	}, "")
}
