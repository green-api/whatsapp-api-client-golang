package main

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/green-api/whatsapp-api-client-golang/pkg/api"
)

func main() {
	GreenAPI := api.GreenAPI{
		IDInstance:       "1101000001",
		APITokenInstance: "d75b3a66374942c5b3c019c698abc2067e151558acbd412345",
	}

	uploadFileResponse, err := GreenAPI.Methods().Sending().UploadFile("example.png")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(uploadFileResponse)

	urlFile := uploadFileResponse["urlFile"].(string)
	request, _ := http.NewRequest("GET", urlFile, nil)
	fileName := path.Base(request.URL.Path)

	sendFileByUrlResponse, err := GreenAPI.Methods().Sending().SendFileByUrl(map[string]interface{}{
		"chatId":   "11001234567@c.us",
		"urlFile":  urlFile,
		"fileName": fileName,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sendFileByUrlResponse)
}
