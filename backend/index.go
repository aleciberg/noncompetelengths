// Go package
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type State struct {
	id         int    `json:"stateID"`
	name       string `json:"name"`
	length     int    `json:"length"`
	law        string `json:"law"`
	lawcode    string `json:"lawcode"`
	date       string `json:"date"`
	experttake string `json:"expertake"`
}

// this will be useful as is
type JsonResponse struct {
	Type    string  `json:"type"`
	Data    []State `json:"data"`
	Message string  `json:"message"`
}

// TO BE AWARE OF
// GOLANG DOES NOT HOT RELOAD
// YOU MUST RESTART SERVER TO SEE CHANGES

// Main function
func main() {
	// Init the mux router
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler)

	// Get all states
	router.HandleFunc("/states", GetStates).Methods("GET")

	// to handle post requests
	router.HandleFunc("/create", CreateState).Methods("POST")

	// serve the app
	fmt.Println("Server at 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "password"
	DB_NAME     = "StateInfoNonCompete"
	// I THINK THE PASSWORD IS password
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	printMessage("Hitting Home Endpoint")
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func GetStates(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Getting states...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM states")

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var states []State

	// Foreach movie
	for rows.Next() {
		var id int
		var name string
		var length int
		var law string
		var lawcode string
		var date string
		var experttake string

		err = rows.Scan(&id, &name, &length, &law, &lawcode, &date, &experttake)

		// check errors
		checkErr(err)

		states = append(states, State{name: name, id: id, length: length, law: law, lawcode: lawcode, date: date, experttake: experttake})
	}

	var response = JsonResponse{Type: "success", Data: states}

	json.NewEncoder(w).Encode(response)
}

// DB set up
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		checkErr(err)
	}

	return db
}

func CreateState(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	name := r.FormValue("name")
	length := r.FormValue("length")
	law := r.FormValue("law")
	lawcode := r.FormValue("lawcode")
	date := r.FormValue("date")
	experttake := r.FormValue("experttake")

	var response = JsonResponse{}

	if id == "" || name == "" {
		response = JsonResponse{Type: "error", Message: "You are missing a parameter."}
	} else {
		db := setupDB()

		printMessage("Inserting state into DB")

		fmt.Println("Inserting new movie with ID: " + id + " and name: " + name)

		var lastInsertID int
		err := db.QueryRow("INSERT INTO states(id, name, length, law, lawcode, date, experttake) VALUES($1, $2, $3, $4, $5, $6, $7) returning id;", id, name, length, law, lawcode, date, experttake).Scan(&lastInsertID)

		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The state has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

func UpdateState(w http.ResponseWriter, r *http.Request) {
	// to update a state in the database
}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}
