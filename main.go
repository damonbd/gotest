package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	//local
	"rolljimmy/controllers"
)

var tpl *template.Template
var router *mux.Router

//probably move to another file
var fm = template.FuncMap{
	"uc": strings.ToUpper,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*"))
	router = mux.NewRouter()
}

func main() {
	addControllerRoutes()
	log.Fatal(http.ListenAndServe(":8000", router))
}

//addControllerRoutes ...
func addControllerRoutes() {
	router = mux.NewRouter()
	personcontroller.AddRoutes(router, tpl)
}
