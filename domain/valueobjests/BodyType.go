package valueobjests

import "errors"

type BodyType struct {
	value string
}

var allowedTypes = map[string]bool{
	"json": true,
}

func NewReqType(val string) (*BodyType, error) {
	if !allowedTypes[val] {
		return nil, errors.New("currently Unsupported Type")
	}
	return &BodyType{value: val}, nil
}
