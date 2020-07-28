package core

import (
	"bytes"
	"io"
	"net/http"
)

func SimpleRequest(client *http.Client, url string, method string, header map[string]string, body []byte) (*http.Response, error) {
	if client == nil {
		client = http.DefaultClient
	}
	var requestBody io.Reader
	if body != nil {
		requestBody = bytes.NewReader(body)
	}
	request, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return nil, err
	}
	if header != nil {
		for k, v := range header {
			request.Header.Add(k, v)
		}
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
