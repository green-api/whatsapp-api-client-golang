package methods

type ReceivingCategory struct {
	GreenAPI GreenAPIInterface
}

func (c ReceivingCategory) ReceiveNotification() (map[string]interface{}, error) {
	return c.GreenAPI.Request("GET", "ReceiveNotification", nil, "")
}

func (c ReceivingCategory) DeleteNotification(receiptId int) (map[string]interface{}, error) {
	return c.GreenAPI.Request("DELETE", "DeleteNotification", map[string]interface{}{
		"receiptId": receiptId,
	}, "")
}

func (c ReceivingCategory) DownloadFile(chatId, idMessage string) (map[string]interface{}, error) {
	return c.GreenAPI.Request("POST", "downloadFile", map[string]interface{}{
		"chatId":    chatId,
		"idMessage": idMessage,
	}, "")
}
