package httptools

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func SendRequest(url string, method string, data string, accept string, contentType string) (int, string, error) {
	client := &http.Client{}
	request, err := http.NewRequest(method, url, strings.NewReader(data))
	if err != nil {
		return 0, "", err
	}
	request.Header.Add("Accept", accept)
	request.Header.Add("Content-Type", contentType)
	response, err := client.Do(request)
	if err != nil {
		return 0, "", err
	}
	if response == nil {
		return 0, "", nil
	}
	defer func() {
		_ = response.Body.Close()
	}()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, "", err
	}
	return response.StatusCode, string(body), nil
}

