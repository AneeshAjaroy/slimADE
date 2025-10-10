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
	StatusCode int
	Headers    map[string]string
	Cookie     map[string]string
	Body       []byte
}

type Transport struct {
	Timeout  time.Duration
	InSecure bool
}

type Adeo struct {
	ID       string
	Name     string
	Request  Request
	Response Response
}
