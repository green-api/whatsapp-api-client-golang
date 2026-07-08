package main

import (
	"log"

	"github.com/green-api/whatsapp-api-client-golang/pkg/api"
	"github.com/green-api/whatsapp-api-client-golang/pkg/categories/methods"
)

func main() {
	GreenAPI := api.GreenAPI{
		IDInstance:       "1101000001",
		APITokenInstance: "d75b3a66374942c5b3c019c698abc2067e151558acbd412345",
	}

	response, err := GreenAPI.Methods().Sending().ForwardMessages(
		"79876543210@c.us",
		"79876782211@c.us",
		[]string{
			"Hi",
			"How are you?",
			"I like tomato :)",
		},
		methods.OptionalTypingTime(12),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response)
}
