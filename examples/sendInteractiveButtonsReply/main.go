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

	replyButtons := []map[string]any{
		{
			"buttonId":   "1",
			"buttonText": "Reply 1",
		},
		{
			"buttonId":   "2",
			"buttonText": "Reply 2",
		},
	}

	parameters := map[string]any{
		"chatId":  "11001234567@c.us",
		"body":    "Choose an option:",
		"buttons": replyButtons,
	}

	response, err := GreenAPI.Methods().Sending().SendInteractiveButtonsReply(parameters)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response)
}
