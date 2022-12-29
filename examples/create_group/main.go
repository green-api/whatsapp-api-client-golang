package main

import (
	"fmt"
	"os"

	"github.com/green-api/whatsapp-api-client-golang/v1/pkg/api"
)

func main() {
	IDInstance := os.Getenv("ID_INSTANCE")
	APITokenInstance := os.Getenv("API_TOKEN_INSTANCE")

	GreenAPI := api.GreenAPI{
		IDInstance:       IDInstance,
		APITokenInstance: APITokenInstance,
	}

	response := GreenAPI.Methods().Groups().CreateGroup("groupName", []string{
		"79001234567@c.us",
		"79002345678@c.us",
	})

	fmt.Println(response)
}
