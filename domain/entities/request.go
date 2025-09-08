package entities

import "api-tester/domain/valueobjests"

type Request struct {
	Id      string
	Name    valueobjests.Name
	Url     valueobjests.Url
	Verb    valueobjests.ReqVerb
	Headers []valueobjests.Header
}
