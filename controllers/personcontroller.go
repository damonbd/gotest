package personcontroller

import (
	"encoding/json"
	"html/template"
	"net/http"
	"rolljimmy/models"
	"strconv"

	"github.com/gorilla/mux"
)

var tpl *template.Template

var people []models.Person

func init() {
	people = append(people, models.Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &models.Address{City: "City X", State: "State X"}})
	people = append(people, models.Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &models.Address{City: "City Z", State: "State Y"}})
	people = append(people, models.Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
}

func personcontroller() {

}

//AddRoutes ...
func AddRoutes(r *mux.Router, mainTpl *template.Template) {
	tpl = mainTpl

	r = r.PathPrefix("/person").Subrouter()

	r.HandleFunc("/testsubroute", TestSubroute)
	r.HandleFunc("/differentcontroller", DifferentController)
	r.HandleFunc("/people", GetPeople).Methods("GET")
	r.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	r.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	r.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	r.HandleFunc("/KruthSucks", KruthSucks)
	r.HandleFunc("/testHtml", TestHTML)
	r.HandleFunc("/passdata", PassData)
	r.HandleFunc("/passperson", PassPerson)
	r.HandleFunc("/createpersonform", CreatePersonForm)
	r.HandleFunc("/postaperson", PostAPerson)
	r.HandleFunc("/passpeople", PassPeople)
	r.HandleFunc("/uppercasestring", UppercaseString)
}

//DifferentController ...
func DifferentController(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "differentcontroller.html", nil)
}

//TestSubroute ...
func TestSubroute(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "testsubroute.html", nil)
}

//KruthSucks ...
func KruthSucks(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "kruthsucks.gohtml", nil)
}

//TestHTML ...
func TestHTML(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "testHtml.html", nil)
}

//PassData ...
func PassData(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "passdata.html", "this string is data passed from my func main in main.go")
}

//PassPerson ...
func PassPerson(w http.ResponseWriter, r *http.Request) {
	person := people[0]

	tpl.ExecuteTemplate(w, "passperson.html", person)
}

//CreatePersonForm ...
func CreatePersonForm(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "createpersonform.html", nil)
}

//PostAPerson ...
func PostAPerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/people", http.StatusSeeOther)
		return
	}
	x := len(people)
	x = x + 1

	people = append(people, models.Person{ID: strconv.Itoa(x), Firstname: r.FormValue("Firstname"), Lastname: r.FormValue("Lastname")})

	json.NewEncoder(w).Encode(people)
}

//PassPeople ...
func PassPeople(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "passpeople.html", people)
}

//GetPeople does things
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

//UppercaseString ...
func UppercaseString(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "uppercasestring.html", "this text should be changed to uppercase.")
}

//GetPerson does things
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

//CreatePerson does things
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person models.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

//DeletePerson does things
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}
