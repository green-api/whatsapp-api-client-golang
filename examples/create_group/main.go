package main

import (
	"fmt"
	"log"
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

	response, err := GreenAPI.Methods().Groups().CreateGroup("groupName", []string{
		"79001234567@c.us",
		"79002345678@c.us",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}
