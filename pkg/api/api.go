package api

import (
	"strconv"
	"strings"

	"github.com/green-api/whatsapp-api-client-golang/pkg/categories"
)

type GreenAPI struct {
	URL              string
	IDInstance       string
	APITokenInstance string
}

func (a GreenAPI) Methods() categories.GreenAPICategories {
	return categories.GreenAPICategories{GreenAPI: a}
}

func (a GreenAPI) Webhook() GreenAPIWebhook {
	return GreenAPIWebhook{GreenAPI: a}
}

func (a GreenAPI) Request(method, APIMethod string, data map[string]interface{}, filePath string) (map[string]interface{}, error) {
	url := a.getURL(method, APIMethod, data)

	return executeRequest(method, url, data, filePath)
}

func (a GreenAPI) getURL(method, APIMethod string, data map[string]interface{}) string {
	if a.URL != "" {
		return a.URL
	}

	var url strings.Builder

	url.WriteString("https://api.green-api.com/")
	url.WriteString("waInstance")
	url.WriteString(a.IDInstance)
	url.WriteString("/")
	url.WriteString(APIMethod)
	url.WriteString("/")
	url.WriteString(a.APITokenInstance)

	if method == "DELETE" {
		url.WriteString("/")
		url.WriteString(strconv.Itoa(data["receiptId"].(int)))
	}

	return url.String()
}
