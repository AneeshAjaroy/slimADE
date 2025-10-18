package server

import (
	requests "api-tester/internal/Requests"
	"html/template"
	"net/http"
)

func InitRouter(rh *requests.RequestHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /request", rh.RequestPage)
	mux.HandleFunc("POST /request", rh.MakeRequest)

	fs := http.FileServer(http.Dir("web"))
	mux.Handle("/js/", fs)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h, pattern := mux.Handler(r)
		if pattern == "" {
			tmpl, err := template.ParseFiles("web/templates/error.html")
			if err != nil {
				panic(err)
			}
			tmpl.Execute(w, map[string]string{"Error": "The Requested Path does not Exist"})
			return
		}
		h.ServeHTTP(w, r)

	})
	return handler
}
