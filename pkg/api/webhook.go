package api

import "time"

type GreenAPIWebhook struct {
	GreenAPI GreenAPI

	ErrorChannel chan error
}

var running = true

func (w GreenAPIWebhook) Start(handler func(map[string]any)) {
	for running {
		response, err := w.GreenAPI.Methods().Receiving().ReceiveNotification()
		if err != nil {
			w.ErrorChannel <- err

			time.Sleep(time.Second * 5)

			continue
		}

		if response != nil {
			body := response["body"]
			handler(body.(map[string]any))

			receiptId := int(response["receiptId"].(float64))
			_, err = w.GreenAPI.Methods().Receiving().DeleteNotification(receiptId)
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
