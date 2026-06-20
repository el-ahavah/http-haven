/*
Exercise 7: Simple Redirector (Status Codes)
Goal: Create a /legacy route. Whenever a user hits this endpoint, permanently redirect them to a new route /v2 with a friendly "Welcome to version 2" message.
Key Tasks:
Redirect traffic using the http.Redirect helper function.
Use the proper status code for a permanent move (http.StatusMovedPermanently).
*/
package main

import (
	"fmt"
	"net/http"
)

func legacyhandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/v2", http.StatusMovedPermanently)
}

func v2handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to version 2")
}

func main() {
	http.HandleFunc("/legacy", legacyhandler)
	http.HandleFunc("/v2", v2handler)

	fmt.Println("server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
