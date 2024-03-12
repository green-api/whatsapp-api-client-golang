package main

import (
	"fmt"

	"github.com/green-api/whatsapp-api-client-golang/pkg/api"
)

func main() {
	Partner := api.GreenAPI{
		IDInstance:       "",
		APITokenInstance: "Partner-Token", // Partner token
	}

	response, err := Partner.Methods().Partner().GetInstances()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response)
}
