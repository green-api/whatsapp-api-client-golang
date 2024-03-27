package greenapi

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/valyala/fasthttp"
)

type GreenAPI struct {
	Host             string
	MediaHost        string
	IDInstance       string
	APITokenInstance string
	PartnerToken     string
}

// func (a GreenAPI) Webhook() GreenAPIWebhook {
// 	return GreenAPIWebhook{
// 		GreenAPI: a,

// 		ErrorChannel: make(chan error),
// 	}
// }

func (a GreenAPI) Request(httpMethod, APImethod string, data map[string]interface{}) (interface{}, error) {
	client := &fasthttp.Client{}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	if APImethod == "createInstance" || APImethod == "deleteInstanceAccount" || APImethod == "getInstances" {
		req.SetRequestURI(fmt.Sprintf("%s/partner/%s/%s", a.Host, APImethod, a.PartnerToken))
	} else {
		req.SetRequestURI(fmt.Sprintf("%s/waInstance%s/%s/%s", a.Host, a.IDInstance, APImethod, a.APITokenInstance))
	}

	req.Header.SetMethod(httpMethod)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("ошибка при сериализации данных в JSON: %s", err)
	}
	req.SetBody([]byte(jsonData))

	// TODO: handle fileUpload cases

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := client.Do(req, resp); err != nil {
		return nil, fmt.Errorf("ошибка при запроса: %s", err)
	}

	var response interface{}

	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshal byte response: %s", err)
	}

	return response, nil
}

func (a GreenAPI) ArrayRequest(url, method string, data map[string]interface{}, filePath string) ([]interface{}, error) {
	response, err := executeRequest(method, url, data, filePath)

	return response.([]interface{}), err
}

func (a GreenAPI) GetURL(method, APIMethod string, data map[string]interface{}) string {
	if a.Host != "" {
		return a.Host
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

func (a GreenAPI) GetPartnerURL(APIMethod string) (string, error) {
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
