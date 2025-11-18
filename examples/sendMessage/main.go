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

	response, err := GreenAPI.Methods().Sending().SendMessage(map[string]any{
		"chatId":     "11001234567@c.us",
		"message":    "Any message",
		"typingTime": 5000,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response)
}
