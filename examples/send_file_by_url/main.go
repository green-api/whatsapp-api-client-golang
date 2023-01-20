package main

import (
	"fmt"
	"log"
	//"os"

	"github.com/green-api/whatsapp-api-client-golang/pkg/api"
)

func main() {
	//You can set environment variables in your OS
	//
	//IDInstance := os.Getenv("ID_INSTANCE")
	//APITokenInstance := os.Getenv("API_TOKEN_INSTANCE")

	GreenAPI := api.GreenAPI{
		IDInstance:       "IDInstance",
		APITokenInstance: "APITokenInstance",
	}

	response, err := GreenAPI.Methods().Sending().SendFileByUrl(map[string]interface{}{
		"chatId":   "11001234567@c.us",
		"urlFile":  "https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Blue.svg",
		"fileName": "Go-Logo_Blue.svg",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}
