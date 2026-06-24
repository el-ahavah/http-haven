/*
Exercise 1: Basic Ping-Pong Server
Goal: Build a minimal web server that listens on port 8080 and responds with "pong" when a user visits the /ping route.
Tasks:
Create a route handler for /ping using http.HandleFunc.
Use w.Write() or fmt.Fprint() to send a plain text response.
Start the server on port :8080 using http.ListenAndServe.
*/
package main

import (
	"fmt"
	"net/http"
)

func pinghandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}
