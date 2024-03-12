package main

import (
	"fmt"

	"github.com/green-api/whatsapp-api-client-golang/pkg/api"
)

func main() {
	Partner := api.GreenAPI{
		IDInstance:       "",
		APITokenInstance: "PARTNERTOKEN", // Partner token
	}

	response, err := Partner.Methods().Partner().DeleteInstanceAccout(1101915618)
	if err != nil {
		fmt.Println(err)
	}

	for key, value := range response {
		fmt.Printf("%s: %v\n", key, value)
	}
}
