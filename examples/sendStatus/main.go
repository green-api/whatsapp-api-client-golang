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

	response, err := GreenAPI.Methods().Status().SendTextStatus(map[string]interface{}{
		"message":         "I sent this status using Green Api Go SDK!",
		"backgroundColor": "#87CEEB",
		"font":            "SERIF",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}
