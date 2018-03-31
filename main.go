package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"rolljimmy/models"
	"strconv"

	"github.com/gorilla/mux"
)

var tpl *template.Template

var testy models.Address

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	testy.City = "test"
	testy.State = "fuck"

	people = append(people, models.Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &models.Address{City: "City X", State: "State X"}})
	people = append(people, models.Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &models.Address{City: "City Z", State: "State Y"}})
	people = append(people, models.Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

	router := mux.NewRouter()

	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	router.HandleFunc("/KruthSucks", KruthSucks)
	router.HandleFunc("/testHtml", TestHtml)
	router.HandleFunc("/passdata", PassData)
	router.HandleFunc("/passperson", PassPerson)
	router.HandleFunc("/createpersonform", CreatePersonForm)
	router.HandleFunc("/postaperson", PostAPerson)
	router.HandleFunc("/passpeople", PassPeople)

	log.Fatal(http.ListenAndServe(":8000", router))
}

//KruthSucks ...
func KruthSucks(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "kruthsucks.gohtml", nil)
}

//TestHtml ...
func TestHtml(w http.ResponseWriter, r *http.Request) {
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

// //Person is a person
// type Person struct {
// 	ID        string   `json:"id,omitempty`
// 	Firstname string   `json:"firstname,omitempty"`
// 	Lastname  string   `json:"lastname,omitempty"`
// 	Address   *Address `json:"address,omitempty"`
// }

var people []models.Person

// //Address is an address
// type Address struct {
// 	City  string `json:"city,omitempty"`
// 	State string `json:"state,omitempty"`
// }
