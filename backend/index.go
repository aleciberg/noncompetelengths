// Go package
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// yeah i will need to switch this
type Movie struct {
	MovieID   string `json:"movieid"`
	MovieName string `json:"moviename"`
}

// this will be useful as is
type JsonResponse struct {
	Type    string  `json:"type"`
	Data    []Movie `json:"data"` // except this will not be movie
	Message string  `json:"message"`
}

// Main function
func main() {
	// Init the mux router
	router := mux.NewRouter()

	// Route handles & endpoints

	// Get all movies
	// router.HandleFunc("/movies/", GetMovies).Methods("GET")

	// need to build to get specific state

	// serve the app
	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8000", router))
}

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "12345678"
	DB_NAME     = "movies"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
    db := setupDB()

    printMessage("Getting movies...")

    // Get all movies from movies table that don't have movieID = "1"
    rows, err := db.Query("SELECT * FROM movies")

    // check errors
    checkErr(err)

    // var response []JsonResponse
    var movies []Movie

    // Foreach movie
    for rows.Next() {
        var id int
        var movieID string
        var movieName string

        err = rows.Scan(&id, &movieID, &movieName)

        // check errors
        checkErr(err)

        movies = append(movies, Movie{MovieID: movieID, MovieName: movieName})
    }

    var response = JsonResponse{Type: "success", Data: movies}

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

func checkErr("something is wrong" err error) {
	printMessage(err)
}

func printMessage(message string, err error) {
    fmt.Println("")
    fmt.Println(message, err)
    fmt.Println("")
}