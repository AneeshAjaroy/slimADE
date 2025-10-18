package requests

import (
	"net/http"
	"strings"
)

type RequestService struct {
	repository RequestRepository
	client     *http.Client
	Transport  *http.Transport
}

func (rs *RequestService) MakeRequest(url string, method string, headers, queryVals map[string]string, body string) error {
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return err
	}
	res, err := rs.client.Do(req)
	if err != nil {
		return err
	}
}
