package main

import (
	"fmt"
	"log"

	"github.com/green-api/whatsapp-api-client-golang/pkg/api"
)

func main() {
	GreenAPI := api.GreenAPI{
		IDInstance:       "1101000001",
		APITokenInstance: "d75b3a66374942c5b3c019c698abc2067e151558acbd412345",
	}

	buttons := []map[string]interface{}{
		{
			"type":       "copy",
			"buttonId":   "1",
			"buttonText": "Copy me",
			"copyCode":   "3333",
		},
		{
			"type":        "call",
			"buttonId":    "2",
			"buttonText":  "Call me",
			"phoneNumber": "79123456789",
		},
		{
			"type":       "url",
			"buttonId":   "3",
			"buttonText": "Green-api",
			"url":        "https://green-api.com",
		},
	}

	parameters := map[string]interface{}{
		"chatId":  "11001234567@c.us",
		"body":    "Main message text",
		"header":  "Message header",
		"footer":  "Message footer",
		"buttons": buttons,
	}
	response, err := GreenAPI.Methods().Sending().SendInteractiveButtons(parameters)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}
