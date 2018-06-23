package main

import (
	// "encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/guild", GetGuild).Methods("GET")
	router.HandleFunc("/guild/feed", GetFeed).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))

}

// GetGuild fetches genetic guild data and returns it
func GetGuild(w http.ResponseWriter, r *http.Request) {

}

// GetFeed returns the activity feed of the guild
func GetFeed(w http.ResponseWriter, r *http.Request) {

}
