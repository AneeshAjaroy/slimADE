package api

import "net/http"

func Init() {
	mux := http.NewServeMux()

	mux.Handle("/")
}
