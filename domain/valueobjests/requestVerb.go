package valueobjests

import "errors"

type ReqVerb struct {
	value string
}

var allowedVerbs = map[string]bool{
	"POST":    true,
	"GET":     true,
	"PATCH":   true,
	"DELETE":  true,
	"PUT":     true,
	"HEAD":    true,
	"OPTIONS": true,
	"TRACE":   true,
	"CONNECT": true,
}

func NewReqVerb(v string) (*ReqVerb, error) {
	if allowedVerbs[v] {
		return &ReqVerb{value: v}, nil
	}
	return nil, errors.New("not valid verb")
}
