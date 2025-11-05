package requests

import "time"

type Request struct {
	Url     string
	Verb    string
	Headers map[string]string
	Queries map[string]string
	Body    []byte
}

type Response struct {
	StatusCode string
	Color      string
	Headers    map[string]string
	Cookies    map[string]string
	Body       string
}

type Transport struct {
	Timeout  time.Duration
	InSecure bool
}

type RequestObject struct {
	ID       string
	Name     string
	Request  Request
	Response Response
}
