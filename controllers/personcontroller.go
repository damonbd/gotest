package personcontroller

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var tpl *template.Template

var Router *mux.Router

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
	Router = mux.NewRouter()
	//Router.HandleFunc("/differentcontroller", DifferentController)

	Subrouter := Router.PathPrefix("/subroute").Subrouter()
	Subrouter.HandleFunc("/testsubroute", testsubroute)
}

func personcontroller() {

}

//AddRoutes ...
func AddRoutes(r *mux.Router) {
	r.HandleFunc("/differentcontroller", DifferentController)
}

//DifferentController ...
func DifferentController(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "differentcontroller.html", nil)
}

func testsubroute(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "testsubroute.html", nil)
}

//Test ...
func Test() {
	x := 2
	x = x + 1
}
