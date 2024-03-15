package main

import (
	"fmt"

	"github.com/green-api/whatsapp-api-client-golang/pkg/api"
)

func main() {
	Partner := api.GreenAPI{
		PartnerToken: "gac.37ea41ed00d74bc7a0899215312fed55bfd9bcd03a1e48",
	}

	response, err := Partner.Methods().Partner().CreateInstance(map[string]interface{}{
		"stateWebhook":    "yes",
		"incomingWebhook": "yes",
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response)
}
