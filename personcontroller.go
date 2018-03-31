package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	//tpl = template.Must(template.ParseGlob("../templates/*"))
	router := mux.NewRouter()
	router.HandleFunc("/differentcontroller", DifferentController).Methods("GET")
}

func personcontroller() {

}

//DifferentController ...
func DifferentController(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "differentcontroller.html", nil)
}
