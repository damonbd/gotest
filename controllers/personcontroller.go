package personcontroller

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

func personcontroller() {

}

//AddRoutes ...
func AddRoutes(r *mux.Router) {
	r.HandleFunc("/differentcontroller", DifferentController)
}

//AddSubRoutes ...
func AddSubRoutes(r *mux.Router) {
	r = r.PathPrefix("/subroute").Subrouter()
	r.HandleFunc("/testsubroute", testsubroute)
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
