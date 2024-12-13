package main

import (
	"json_responses_requests/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			handlers.HandlePost(w, r)
		case "GET":
			handlers.HandleGet(w, r)
		default:
			handlers.HandleMethodNotAllowed(w, r)
		}
	})
	http.ListenAndServe(":8080", nil)
}
