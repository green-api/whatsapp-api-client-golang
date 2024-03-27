package main

import (
	"fmt"
	"log"

	greenapi "github.com/green-api/whatsapp-api-client-golang"
)

func main() {
	Gabit := greenapi.GreenAPI{
		Host:             "https://api.green-api.com",
		IDInstance:       "",
		APITokenInstance: "",
		PartnerToken:     "",
	}

	response, err := Gabit.Account().GetInstances()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}
