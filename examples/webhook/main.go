package main

import (
	"fmt"
	//"os"

	"github.com/green-api/whatsapp-api-client-golang/pkg/api"
	"github.com/green-api/whatsapp-api-client-golang/pkg/webhook"
)

func main() {
	//You can set environment variables in your OS
	//
	//IDInstance := os.Getenv("ID_INSTANCE")
	//APITokenInstance := os.Getenv("API_TOKEN_INSTANCE")

	GreenAPI := api.GreenAPI{
		IDInstance:       "IDInstance",
		APITokenInstance: "APITokenInstance",
	}

	GreenAPIWebhook := webhook.GreenAPIWebhook{
		GreenAPI: GreenAPI,
	}

	GreenAPIWebhook.Start(func(body map[string]interface{}) {
		fmt.Println(body)
	})
}
