/*
Exercise 5: User-Agent Echo (Reading Headers)
Goal: Create an /agent route that reads the incoming browser or client header details and echoes it back in plain text: "You are visiting us using: [User-Agent Info]".
Key Tasks:
Inspect request headers using r.Header.Get("User-Agent").
Handle instances where the header might be missing or empty.
*/
package main

import (
	"fmt"
	"net/http"
)

func agenthandler(w http.ResponseWriter, r *http.Request) {
	userAgent := r.Header.Get("User-Agent")
	if userAgent == "" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "You are visiting us using: ", userAgent)
}

func main() {
	http.HandleFunc("/agent", agenthandler)
	fmt.Println("server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
