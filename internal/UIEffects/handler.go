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

	tmpl, err = template.ParseFiles("web/templates/UIEffects/header.html")
	if err != nil {
		panic(err)
	}
	pages["headerAdd"] = tmpl

	tmpl, err = template.ParseFiles("web/templates/UIEffects/auth/apikey.html")
	if err != nil {
		panic(err)
	}
	pages["apikey"] = tmpl

	tmpl, err = template.ParseFiles("web/templates/UIEffects/auth/bearer.html")
	if err != nil {
		panic(err)
	}
	pages["bearerAuth"] = tmpl

	tmpl, err = template.ParseFiles("web/templates/error.html")
	if err != nil {
		panic(err)
	}
	pages["errPage"] = tmpl

	return &UIEffectsHandler{
		tmplPages: pages,
	}

}

func (ui *UIEffectsHandler) QueryAdd(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	ui.tmplPages["queryAdd"].Execute(w, nil)
}

func (ui *UIEffectsHandler) QueryRemove(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (ui *UIEffectsHandler) HeaderAdd(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	ui.tmplPages["headerAdd"].Execute(w, nil)
}

func (ui *UIEffectsHandler) HeaderRemove(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (ui *UIEffectsHandler) AuthPage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(500)
		ui.tmplPages["errPage"].Execute(w, map[string]string{"Error": err.Error()})
		return
	}
	if r.PostFormValue("auth") == "Bearer Token" {
		w.WriteHeader(200)
		ui.tmplPages["bearerAuth"].Execute(w, nil)
		return
	}

	if r.PostFormValue("auth") == "apiKey" {
		w.WriteHeader(200)
		ui.tmplPages["apikey"].Execute(w, nil)
		return
	}
}

func (ui *UIEffectsHandler) BodyType(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(500)
		ui.tmplPages["errPage"].Execute(w, map[string]string{"Error": err.Error()})
		return
	}
	if r.PostFormValue("auth") == "Bearer Token" {
		w.WriteHeader(200)
		ui.tmplPages["bearerAuth"].Execute(w, nil)
		return
	}

	if r.PostFormValue("auth") == "apiKey" {
		w.WriteHeader(200)
		ui.tmplPages["apikey"].Execute(w, nil)
		return
	}
}
