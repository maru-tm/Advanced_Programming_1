package main

import (
	"json_responses_requests/handlers"
	"json_responses_requests/storage"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Call ConnectDatabase to initialize the DB connection
	storage.ConnectDatabase()

	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./template")))

	// Routes for CRUD operations
	r.HandleFunc("/user", handlers.CreateUserHandler).Methods("POST")
	r.HandleFunc("/user/{id:[0-9]+}", handlers.GetUserHandler).Methods("GET")
	r.HandleFunc("/user/{id:[0-9]+}", handlers.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/user/{id:[0-9]+}", handlers.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/users", handlers.GetAllUsersHandler).Methods("GET")

	// Static route for HTML page

	log.Fatal(http.ListenAndServe(":8080", r))
	log.Println("Starting server on :8080...")
}
