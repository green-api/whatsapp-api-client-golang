package main

import (
	"log"

	"github.com/green-api/whatsapp-api-client-golang/pkg/api"
)

func main() {
	GreenAPI := api.GreenAPI{
		IDInstance:       "1101000001",
		APITokenInstance: "d75b3a66374942c5b3c019c698abc2067e151558acbd412345",
	}

	response, err := GreenAPI.Methods().Service().EditMessage(
		map[string]any{
			"chatId":    "79876543210@c.us",
			"idMessage": "21312312",
			"message":   "Hello World!",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response)
}
