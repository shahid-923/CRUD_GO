package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:firstname`
	Lastname  string `json:lastname`
}

var movies []Movie

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") 
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)   //Decode the JSON request body into a Movie struct
	movie.ID = strconv.Itoa(rand.Intn(10000000)) //Generate a random ID for the new movie
	movies = append(movies, movie)               //Append the new movie to the movies slice(main data store)
	json.NewEncoder(w).Encode(movie)             //Encode the newly created movie to JSON and send it as the response

}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") 
	json.NewEncoder(w).Encode(movies)                           
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //The response must be sent in JSON format → so you must encode again.
	params := mux.Vars(r)                              //extracts the URL path parameters (like {id}) from the incoming request.
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) //Remove the movie from the slice by appending the elements before and after the index.
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]          //Set the ID of the updated movie to the same ID as the original movie.
			movies = append(movies, movie)   //Append the updated movie to the movies slice.
			json.NewEncoder(w).Encode(movie) //Encode the updated movie to JSON and send it as the response.
			return
		}
	}
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //The response must be sent in JSON format → so you must encode again.
	params := mux.Vars(r)                              //extracts the URL path parameters (like {id}) from the incoming request.
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) //Remove the movie from the slice by appending the elements before and after the index.
			break
		}
	}
	json.NewEncoder(w).Encode(movies) //Encode the updated movies slice to JSON and send it as the response.
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)                            // extracts the URL path parameters (like {id}) from the incoming request.
	for _, item := range movies {  
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
func main() {
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})     // hardcoding
	movies = append(movies, Movie{ID: "2", Isbn: "454555", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})
	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movie", createMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")
	fmt.Println("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
