package api

import (
	"fmt"
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
		GreenAPI:     a,
		ErrorChannel: make(chan error),
	}
}

func (a GreenAPI) Request(method, APIMethod string, data map[string]any, filePath string) (map[string]any, error) {
	url := a.getURL(method, APIMethod, data)
	response, err := executeRequest(method, url, data, filePath)

	if response == nil {
		return nil, err
	}

	responseMap, ok := response.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("unexpected response type: %T, expected map[string]any", response)
	}

	return responseMap, err
}

func (a GreenAPI) PartnerRequest(method, APIMethod string, data map[string]any, filePath string) (map[string]any, error) {
	url, err := a.getPartnerURL(APIMethod)
	if err != nil {
		return nil, err
	}

	response, err := executeRequest(method, url, data, filePath)

	if response == nil {
		return nil, err
	}

	responseMap, ok := response.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("unexpected response type: %T, expected map[string]any", response)
	}

	return responseMap, err
}

func (a GreenAPI) ArrayPartnerRequest(method, APIMethod string, data map[string]any, filePath string) ([]any, error) {
	url, err := a.getPartnerURL(APIMethod)
	if err != nil {
		return nil, err
	}

	response, err := executeRequest(method, url, data, filePath)

	return response.([]any), err
}

func (a GreenAPI) RawRequest(method, APIMethod string, data map[string]any, filePath string) (any, error) {
	url := a.getURL(method, APIMethod, data)

	return executeRequest(method, url, data, filePath)
}

func (a GreenAPI) ArrayRequest(method, APIMethod string, data map[string]any, filePath string) ([]any, error) {
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

	return response.([]any), err
}

func (a GreenAPI) getURL(method, APIMethod string, data map[string]any) string {
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

func (a GreenAPI) getPartnerURL(APIMethod string) (string, error) {
	if a.PartnerToken == "" {
		return "", fmt.Errorf("error while generating URL: PartnerToken is empty")
	}

	var url strings.Builder

	url.WriteString("https://api.green-api.com")

	url.WriteString("/")
	url.WriteString("partner")
	url.WriteString("/")
	url.WriteString(APIMethod)
	url.WriteString("/")
	url.WriteString(a.PartnerToken)

	return url.String(), nil
}
