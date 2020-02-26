package httptools

import (
	"net/http"
	"strings"
)

func SendRequest(url string, method string, data string, accept string, contentType string) (*http.Response, error) {
	client := &http.Client{}
	request, err := http.NewRequest(method, url, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", accept)
	request.Header.Add("Content-Type", contentType)
	response, err := client.Do(request)
	if err != nil {
		return response, err
	}
	return response,nil
}

