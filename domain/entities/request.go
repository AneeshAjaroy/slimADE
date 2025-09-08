package entities

import "api-tester/domain/valueobjests"

type Request struct {
	Id              string
	Name            valueobjests.Name
	Url             valueobjests.Url
	Verb            valueobjests.ReqVerb
	RequestHeaders  []valueobjests.Header
	ResponseHeaders []valueobjests.Header
	RequestType     valueobjests.BodyType
	ReaponseType    valueobjests.BodyType
	RequestBody     []byte
	ResponseBody    []byte
}
