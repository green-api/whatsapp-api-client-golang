package main

import (
	"fmt"
	"log"
	//"os"

	"github.com/green-api/whatsapp-api-client-golang/pkg/api"
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

	response, err := GreenAPI.Methods().Sending().SendFileByUpload("example.png", map[string]interface{}{
		"chatId": "11001234567@c.us",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}