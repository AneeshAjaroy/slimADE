package server

import (
	"html/template"
	"net/http"
)

type Handler struct {
	routes map[string]map[string]func(http.ResponseWriter, *http.Request)
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, ok := h.routes[r.Method]; ok {
		handler(w, r)
		return
	}
	tmpl, err := template.ParseFiles("../web/templates/NotAllowed.html")
	if err != nil {
		panic("Static Files Missing")
	}
	tmpl.Execute(w, nil)

}

func (h Handler) RegisterRoutes(method string, path string, handler func(http.ResponseWriter, *http.Request)) {
	if _, ok := h.routes[path]; ok {
		h.routes[path][method] = handler
		return
	}
	h.routes[path] = make(map[string]func(http.ResponseWriter, *http.Request))
	h.routes[path][method] = handler
}

func InitRouter() {

}
