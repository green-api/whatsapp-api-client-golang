package main

import (
	"fmt"
	"log"

	greenapi "github.com/green-api/whatsapp-api-client-golang"
)

func main() {
	Gabit := greenapi.GreenAPI{
		Host:             "https://api.green-api.com",
		IDInstance:       "1101912410",
		APITokenInstance: "0bcb6214e0374d4a9c54c73dba0f6cb2807468c3f57c41e08b",
	}

	response, err := Gabit.Account().GetSettings()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}
