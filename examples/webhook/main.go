package main

import (
	"fmt"
	"log"
	"os"

	"github.com/green-api/whatsapp-api-client-golang/pkg/api"
	"github.com/green-api/whatsapp-api-client-golang/pkg/webhook"
)

func main() {
	IDInstance := os.Getenv("ID_INSTANCE")
	APITokenInstance := os.Getenv("API_TOKEN_INSTANCE")

	GreenAPI := api.GreenAPI{
		IDInstance:       IDInstance,
		APITokenInstance: APITokenInstance,
	}

	GreenAPIWebhook := webhook.GreenAPIWebhook{
		GreenAPI: GreenAPI,
	}

	GreenAPIWebhook.Start(func(body map[string]interface{}) {
		typeWebhook := body["typeWebhook"]
		if typeWebhook == "incomingMessageReceived" {
			senderData := body["senderData"]
			chatId := senderData.(map[string]interface{})["chatId"]

			response, err := GreenAPI.Methods().Sending().SendMessage(map[string]interface{}{
				"chatId":  chatId,
				"message": "Any message",
			})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(response)

			GreenAPIWebhook.Stop()
		}
	})
}
