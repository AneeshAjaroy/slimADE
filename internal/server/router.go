package server

import (
	requests "api-tester/internal/Requests"
	uieffects "api-tester/internal/UIEffects"
	"html/template"
	"mime"
	"net/http"
)

func InitRouter(rh *requests.RequestHandler, uh *uieffects.UIEffectsHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /request", rh.RequestPage)
	mux.HandleFunc("POST /request", rh.MakeRequest)

	mux.HandleFunc("GET /ui/queryAdd", uh.QueryAdd)

	mime.AddExtensionType(".css", "text/css")
	fs := http.FileServer(http.Dir("web"))
	mux.Handle("/web/", http.StripPrefix("/web/", fs))

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
