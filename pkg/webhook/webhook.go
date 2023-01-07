package webhook

import (
	"log"

	"github.com/green-api/whatsapp-api-client-golang/pkg/api"
)

type GreenAPIWebhook struct {
	GreenAPI api.GreenAPI
}

var running = true

func (w GreenAPIWebhook) Start(handler func(map[string]interface{})) {
	for running {
		response, err := w.GreenAPI.Methods().Receiving().ReceiveNotification()
		if err != nil {
			log.Fatal(err)
		}

		if response == nil {
			continue
		} else {
			body := response["body"]
			handler(body.(map[string]interface{}))

			receiptId := int(response["receiptId"].(float64))
			response, err = w.GreenAPI.Methods().Receiving().DeleteNotification(receiptId)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func (w GreenAPIWebhook) Stop() {
	running = false
}
