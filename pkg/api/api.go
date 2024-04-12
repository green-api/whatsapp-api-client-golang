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
	PartnerToken     string
}

func (a GreenAPI) Methods() categories.GreenAPICategories {
	return categories.GreenAPICategories{GreenAPI: a}
}

func (a GreenAPI) Webhook() GreenAPIWebhook {
	return GreenAPIWebhook{
		GreenAPI: a,

		ErrorChannel: make(chan error),
	}
}

func (a GreenAPI) Request(method, APIMethod string, data map[string]interface{}, filePath string) (map[string]interface{}, error) {
	url := a.getURL(method, APIMethod, data)
	if APIMethod == "deleteStatus" {
		_, err := executeRequest(method, url, data, filePath)
		if err != nil {
			if err.Error() == "unexpected end of JSON input" {
				return nil, nil
			}
			return nil, err
		}
		return nil, nil
	}

	response, err := executeRequest(method, url, data, filePath)

	return response.(map[string]interface{}), err
}

func (a GreenAPI) PartnerRequest(method, APIMethod string, data map[string]interface{}, filePath string) (map[string]interface{}, error) {
	url := a.getPartnerURL(method, APIMethod, data)

	response, err := executeRequest(method, url, data, filePath)

	return response.(map[string]interface{}), err
}

func (a GreenAPI) ArrayPartnerRequest(method, APIMethod string, data map[string]interface{}, filePath string) ([]interface{}, error) {
	url := a.getPartnerURL(method, APIMethod, data)

	response, err := executeRequest(method, url, data, filePath)

	return response.([]interface{}), err
}

func (a GreenAPI) RawRequest(method, APIMethod string, data map[string]interface{}, filePath string) (interface{}, error) {
	url := a.getURL(method, APIMethod, data)

	return executeRequest(method, url, data, filePath)
}

func (a GreenAPI) ArrayRequest(method, APIMethod string, data map[string]interface{}, filePath string) ([]interface{}, error) {
	url := a.getURL(method, APIMethod, data)
	if APIMethod == "getOutgoingStatuses" || APIMethod == "getIncomingStatuses" {
		if data["minutes"] != nil {
			url = (url + "?minutes=" + data["minutes"].(string))
		}
	}
	if APIMethod == "getStatusStatistic" {
		if data["idMessage"] != nil {
			url = (url + "?idMessage=" + data["idMessage"].(string))
		}
	}
	response, err := executeRequest(method, url, data, filePath)

	return response.([]interface{}), err
}

func (a GreenAPI) getURL(method, APIMethod string, data map[string]interface{}) string {
	if a.URL != "" {
		return a.URL
	}

	var url strings.Builder

	if APIMethod == "SendFileByUpload" || APIMethod == "UploadFile" {
		url.WriteString("https://media.green-api.com")
	} else {
		url.WriteString("https://api.green-api.com")
	}

	url.WriteString("/")
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

func (a GreenAPI) getPartnerURL(method, APIMethod string, data map[string]interface{}) string {
	if a.URL != "" {
		return a.URL
	}

	var url strings.Builder

	url.WriteString("https://api.green-api.com")

	url.WriteString("/")
	url.WriteString("partner")
	url.WriteString("/")
	url.WriteString(APIMethod)
	url.WriteString("/")
	url.WriteString(a.APITokenInstance)

	return url.String()
}
