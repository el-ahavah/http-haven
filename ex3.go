/*
Exercise 3: Text Counter (URL Variables & Methods)
Goal: Build a server with a /count route. If a user sends a GET request, return the text "Send a POST request with text to count words". If they send a POST request, read the text body and return the number of characters.
Key Tasks:
Differentiate between GET and POST methods using r.Method.
Read the entire request body using io.ReadAll(r.Body).
Return the character length as a string.
*/
package main

import (
	"fmt"
	"io"
	"net/http"
)

func counthandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, "Send a POST request with text to count words")
		return
	}

	if r.Method == http.MethodPost {
		text, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		count := len(text)

		fmt.Fprint(w, count, "\n")

	} else {
		http.Error(w, "405 Bad Request", http.StatusBadRequest)
		return
	}
}

func main() {
	http.HandleFunc("/count", counthandler)
	fmt.Println("server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

/* usage:
curl http://localhost:8080/count

curl -X POST http://localhost:8080/count -d "Hello World"
*/
