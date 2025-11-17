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

	response, err := GreenAPI.Methods().Service().SendTyping("11001234567@c.us")
	if err != nil {
		log.Fatalf("SendTyping failed: %v", err)
	}

	fmt.Printf("SendTyping successful: %+v\n", response)
}
