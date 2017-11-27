package service

import (
	"net/http"
	"github.com/unrolled/render"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"os"
)

func GetServer() *negroni.Negroni {
	server := negroni.Classic()
	m := mux.NewRouter()
	formatter := render.New(render.Options {
		IndentJSON: true,
		Extensions: []string{".tmpl", ".html"},
		Directory: "templates",
	})

	initRouters(m, formatter)

	server.UseHandler(m)
	return server
}

func initRouters(m *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
    if len(webRoot) == 0 {
        webRoot = "./static"
	}

	m.HandleFunc("/", indexHandlerFunc(formatter)).Methods("GET")
	m.HandleFunc("/json", apiHandlerFunc(formatter)).Methods("GET")
	m.HandleFunc("/unknown", NotImplement)
	m.HandleFunc("/login", formHandler(formatter)).Methods("POST")
	m.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot))))
}

func indexHandlerFunc(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "login", "")
	}
}

func apiHandlerFunc(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			ID      string `json:"id"`
			Name string `json:"content"`
		} {ID: "8675309", Name: "tangxzh"})
	}
}

func NotImplement(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "501 not implemented.", 501)
}

func formHandler(formatter *render.Render) http.HandlerFunc {
	return func (w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		id := req.FormValue("id")
		name := req.FormValue("name")
		formatter.HTML(w, http.StatusOK, "detail", struct {
			ID string
			NAME string
		} {ID: id, NAME: name})
	}
}
