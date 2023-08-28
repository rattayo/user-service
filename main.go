package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Note struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var notes = []Note{
	{ID: "1", Title: "Max Friedrichs", Content: "Am Schlossgarten 13, Niederzissen"},
	{ID: "2", Title: "Egon Pfaff", Content: "Forststraße 3, Much"},
}

func main() {
	r := mux.NewRouter()

	// Define your routes and handlers
	r.HandleFunc("/user", GetNotes).Methods("GET")
	r.HandleFunc("/health/ready", IsReady).Methods("GET")
	r.HandleFunc("/health/alive", IsAlive).Methods("GET")

	// Start the server
	port := ":8080"
	fmt.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func GetNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func IsReady(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func IsAlive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
