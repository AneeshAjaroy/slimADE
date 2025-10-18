package requests

import (
	"html/template"
	"net/http"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type RequestHandler struct {
	service   RequestService
	validator validator.Validate
	pages     map[string]*template.Template
}

func NewRequestHandler(svc *RequestService) *RequestHandler {
	reqPage, err := template.ParseFiles("web/test.html")
	if err != nil {
		panic(err)
	}
	errPage, err := template.ParseFiles("web/templates/error.html")
	if err != nil {
		panic(err)
	}
	v := validator.New()
	v.RegisterValidation("headerkey", func(fl validator.FieldLevel) bool {
		re := regexp.MustCompile(`^[A-Za-z0-9-]+$`)
		return re.MatchString(fl.Field().String())
	})

	return &RequestHandler{
		service:   *svc,
		validator: *v,
		pages:     map[string]*template.Template{"reqPage": reqPage, "errPage": errPage},
	}

}

func (rh *RequestHandler) RequestPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	rh.pages["reqPage"].Execute(w, nil)
}

func (rh *RequestHandler) MakeRequest(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		rh.pages["errPage"].Execute(w, map[string]string{"Error": err.Error()})
	}

	// validate URL
	url := r.PostFormValue("url")
	err = rh.validator.Var(url, "http_url")
	if err != nil {
		rh.pages["errPage"].Execute(w, map[string]string{"Error": err.Error()})
	}

	//validate VERB
	method := r.PostFormValue("verb")
	err = rh.validator.Var(url, "oneof:POST GET PUT DELETE PATCH HEAD OPTIONS")
	if err != nil {
		rh.pages["errPage"].Execute(w, map[string]string{"Error": err.Error()})
	}

	//validate Headers
	headers := make(map[string]string)
	for i, v := range r.PostForm["key-h"] {
		err = rh.validator.Var(v, "headerkey")
		if err != nil {
			rh.pages["errPage"].Execute(w, map[string]string{"Error": err.Error()})
		}
		if r.PostForm["enabled-h"][i] == "true" {
			headers[v] = r.PostForm["value-h"][i]
			err = rh.validator.Var(headers[v], "printascii")
			if err != nil {
				rh.pages["errPage"].Execute(w, map[string]string{"Error": err.Error()})
			}
		}
	}

	queryVals := make(map[string]string)
	for i, v := range r.PostForm["key-q"] {
		if r.PostForm["enabled-q"][i] == "true" {
			headers[v] = r.PostForm["value-q"][i]
		}
	}
	body := r.PostFormValue("body")

	// err = rh.service.MakeRequest(url, method, headers, queryVals, body)
	// if err != nil {
	// 	rh.pages["errPage"].Execute(w, map[string]string{"Error": err.Error()})
	// }

}
