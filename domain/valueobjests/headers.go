package valueobjests

import (
	"errors"
	"strings"
)

type Header struct {
	key   string
	value []string
}

func isValid(str string) bool {
	for _, s := range str {
		if (s != 45) && (s < 64) && (s > 91) && (s < 96) && (s > 123) {
			return false
		}
	}
	return true
}

func NewHeader(h map[string]string) (*Header, error) {
	for k, v := range h {
		if !isValid(k) {
			return nil, errors.New("header keys must only contain alphabets and hypen")
		}
		values := strings.Split(v, ";")
		return &Header{key: k, value: values}, nil
	}
	return nil, nil
}
