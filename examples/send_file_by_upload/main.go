package main

import (
	"fmt"
	"os"

	"github.com/green-api/whatsapp-api-client-golang/pkg/api"
)

func main() {
	IDInstance := os.Getenv("ID_INSTANCE")
	APITokenInstance := os.Getenv("API_TOKEN_INSTANCE")

	GreenAPI := api.GreenAPI{
		IDInstance:       IDInstance,
		APITokenInstance: APITokenInstance,
	}

	response := GreenAPI.Methods().Sending().SendFileByUpload("example.png", map[string]interface{}{
		"chatId": "79001234567@c.us",
	})

	fmt.Println(response)
}
