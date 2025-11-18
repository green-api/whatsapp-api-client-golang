package main

import (
	"log"

	"github.com/green-api/whatsapp-api-client-golang/pkg/api"
)

func main() {
	Partner := api.GreenAPI{
		PartnerToken: "gac.1234567891234567891234567891213456789",
	}

	response, err := Partner.Methods().Partner().CreateInstance(map[string]any{
		"stateWebhook":    "yes",
		"incomingWebhook": "yes",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response)
}
