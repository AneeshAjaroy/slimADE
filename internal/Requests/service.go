package requests

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type RequestService struct {
	repository RequestRepository
	client     *http.Client
	Transport  *http.Transport
}

func NewRequestService() *RequestService {
	client := &http.Client{}
	return &RequestService{
		client: client,
	}
}

func (rs *RequestService) MakeRequest(url string, method string, headers, queryVals map[string]string, body string) (*Response, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := rs.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	resBody, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(resBody))
	respHeaders := make(map[string]string)
	for k, v := range resp.Header {
		respHeaders[k] = strings.Join(v, ", ")
	}
	var color string

	switch resp.StatusCode / 100 {
	case 2:
		color = "green"
	case 3:
		color = "blue"
	case 4:
		color = "orange"
	case 5:
		color = "red"
	}

	return &Response{
		StatusCode: resp.Status,
		Color:      color,
		Headers:    respHeaders,
		Cookies:    map[string]string{},
		Body:       string(resBody),
	}, nil

}
