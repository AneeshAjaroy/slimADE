package application

type RequestService interface {
	MakeRequest()
	FetchRequest()
	ListRequests()
}
