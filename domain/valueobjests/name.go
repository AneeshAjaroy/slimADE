package valueobjests

import "errors"

type Name struct {
	value string
}

func NewName(val string) (*Name, error) {
	if len(val) >= 16 {
		return nil, errors.New("name cannot be larger than 15 characters")
	}
	return &Name{value: val}, nil

}
