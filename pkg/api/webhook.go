package api

import "time"

type GreenAPIWebhook struct {
	GreenAPI GreenAPI

	ErrorChannel chan error
}

var running = true

func (w GreenAPIWebhook) Start(handler func(map[string]interface{})) {
	for running {
		response, err := w.GreenAPI.Methods().Receiving().ReceiveNotification()
		if err != nil {
			w.ErrorChannel <- err

			time.Sleep(time.Second * 5)

			continue
		}

		if response == nil {
			continue
		} else {
			body := response["body"]
			handler(body.(map[string]interface{}))

			receiptId := int(response["receiptId"].(float64))
			response, err = w.GreenAPI.Methods().Receiving().DeleteNotification(receiptId)
			if err != nil {
				w.ErrorChannel <- err

				time.Sleep(time.Second * 5)

				continue
			}
		}
	}
}

func (w GreenAPIWebhook) Stop() {
	running = false
}
