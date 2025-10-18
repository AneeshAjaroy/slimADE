package main

import (
	requests "api-tester/internal/Requests"
	uieffects "api-tester/internal/UIEffects"
	"api-tester/internal/server"
	"net/http"
)

func main() {
	svc := &requests.RequestService{}
	rh := requests.NewRequestHandler(svc)
	uh := uieffects.NewUIUIEffectsHandler()
	mux := server.InitRouter(rh, uh)
	http.ListenAndServe(":8080", mux)
}
