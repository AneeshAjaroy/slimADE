package main

import (
	requests "api-tester/internal/Requests"
	"api-tester/internal/server"
	"net/http"
)

func main() {
	rh := &requests.RequestHandler{}
	mux := server.InitRouter(rh)
	http.ListenAndServe(":8080", mux)
}
