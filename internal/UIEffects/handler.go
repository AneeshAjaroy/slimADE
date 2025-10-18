package uieffects

import (
	"html/template"
	"net/http"
)

type UIEffectsHandler struct {
	tmplPages map[string]*template.Template
}

func NewUIUIEffectsHandler() *UIEffectsHandler {
	pages := make(map[string]*template.Template)

	tmpl, err := template.ParseFiles("web/templates/UIEffects/query.html")
	if err != nil {
		panic(err)
	}
	pages["queryAdd"] = tmpl

	return &UIEffectsHandler{
		tmplPages: pages,
	}

}

func (ui *UIEffectsHandler) QueryAdd(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	ui.tmplPages["queryAdd"].Execute(w, nil)
}
