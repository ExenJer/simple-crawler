package client

import (
	"errors"
	"io/ioutil"
	"net/http"
)

var (
	ErrUnknownMethod = errors.New("error fetch: unknown method")
)

func FetchBody(method, url string, client *http.Client) (string, error) {
	switch method {
	case http.MethodGet:
		return getMethodFetchBody(url, client)
	default:
		return "", ErrUnknownMethod
	}
}

func getMethodFetchBody(url string, client *http.Client) (string, error) {
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	return string(bytes), nil
}
