package requests

import (
	"html/template"
	"net/http"
)

type RequestHandler struct {
	service RequestService
}

func (rh *RequestHandler) RequestPage(w http.ResponseWriter, r *http.Request) {

	tmplReq, err := template.ParseFiles("../web/test.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(200)
	tmplReq.Execute(w, nil)

}
