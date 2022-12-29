package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func ExecuteRequest(method, url string, data map[string]interface{}, filePath string) (map[string]interface{}, error) {
	client := &http.Client{}

	req := getRequest(method, url, data, filePath)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return getResponse(resp)
}

func getRequest(method, url string, data map[string]interface{}, filePath string) *http.Request {
	if method == http.MethodGet || method == http.MethodDelete {
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			log.Fatal(err)
		}

		return req
	}

	if filePath == "" {
		buf, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}

		req, err := http.NewRequest(method, url, bytes.NewBuffer(buf))
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Add("Content-Type", "application/json")

		return req
	}

	buffer := &bytes.Buffer{}

	writer := multipart.NewWriter(buffer)

	if data != nil {
		for key, value := range data {
			err := writer.WriteField(key, value.(string))
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		log.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = writer.Close()
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(method, url, buffer)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())

	return req
}

func getResponse(resp *http.Response) (map[string]interface{}, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("StatusCode = %d. Body = %s.", resp.StatusCode, body))
	}

	var data map[string]interface{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}
