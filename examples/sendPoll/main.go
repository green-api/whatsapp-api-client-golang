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

	response, err := GreenAPI.Methods().Sending().SendPoll(map[string]interface{}{
		"chatId":  "11001234567@c.us",
		"message": "Please choose the color:",
		"options": []map[string]interface{}{
			{
				"optionName": "green",
			},
			{
				"optionName": "red",
			},
			{
				"optionName": "blue",
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}
