package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

func executeRequest(method, url string, data map[string]any, filePath string) (any, error) {
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

func getRequest(method, url string, data map[string]any, filePath string) (*http.Request, error) {
	if data == nil || method == http.MethodGet || method == http.MethodDelete {
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

	for key, value := range data {
		err := writer.WriteField(key, value.(string))
		if err != nil {
			return nil, err
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
	req.Header.Set("GA-Filename", filepath.Base(filePath))

	return req, err
}

func getResponse(resp *http.Response) (any, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if len(body) > 0 {
			return nil, fmt.Errorf("StatusCode: %d. Body: %s", resp.StatusCode, strings.TrimSpace(string(body)))
		}
		return nil, fmt.Errorf("StatusCode: %d", resp.StatusCode)
	}

	// amazing
	if strings.TrimSpace(string(body)) == `{"code":401,"description":"Unauthorized"}` {
		return nil, fmt.Errorf("StatusCode: %d. Body: %s", 401, strings.TrimSpace(string(body)))
	}

	if len(body) == 0 {
		return nil, nil
	}

	var data any
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
