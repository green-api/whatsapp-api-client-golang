package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

func executeRequest(method, url string, data map[string]interface{}, filePath string) (map[string]interface{}, error) {
	client := &http.Client{}

	req, err := getRequest(method, url, data, filePath)
	if strings.Contains(url, "uploadFile") {
		req, err = getUploadFileRequest(method, url, filePath)
	}

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return getResponse(resp)
}

func getRequest(method, url string, data map[string]interface{}, filePath string) (*http.Request, error) {
	if method == http.MethodGet || method == http.MethodDelete {
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}

		return req, nil
	}

	if filePath == "" {
		buf, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		req, err := http.NewRequest(method, url, bytes.NewBuffer(buf))
		if err != nil {
			return nil, err
		}

		req.Header.Set("Content-Type", "application/json")

		return req, nil
	}

	buffer := &bytes.Buffer{}

	writer := multipart.NewWriter(buffer)

	if data != nil {
		for key, value := range data {
			err := writer.WriteField(key, value.(string))
			if err != nil {
				return nil, err
			}
		}
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, nil
}

func getUploadFileRequest(method, url string, filePath string) (*http.Request, error) {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	MIMEType := mimetype.Detect(buf).String()

	req.Header.Set("Content-Type", MIMEType)

	return req, nil
}

func getResponse(resp *http.Response) (map[string]interface{}, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("StatusCode = %d. Body = %s.", resp.StatusCode, body))
	}

	var data map[string]interface{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
