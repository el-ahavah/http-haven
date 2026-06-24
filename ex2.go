/*
Exercise 2: Query Parameters & Path Validation
Goal: Create a /hello endpoint that reads a name query parameter (e.g., /hello?name=Alice) and responds with "Hello, Alice!". If the parameter is missing, default to "Hello, Guest!".
Tasks:
Extract query parameters using r.URL.Query().Get("name").
Reject any HTTP method that is not GET by returning an http.StatusMethodNotAllowed status code.
*/
package main

import (
	"fmt"
	"net/http"
)

func hellohandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "Guest"
	}

	fmt.Fprint(w, "Hello,", " ", name, "!")
}
