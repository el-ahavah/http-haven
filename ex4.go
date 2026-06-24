/*
Exercise 4: Basic Math API (Multiple Query Parameters)
Goal: Create a /calculate route that accepts three query parameters: op (operation), a, and b. For example, /calculate?op=add&a=10&b=5 should respond with Result: 15.
Key Tasks:
Parse string query variables using r.URL.Query().Get().
Convert string inputs to integers using strconv.Atoi().
Support add, subtract, and multiply. Return an HTTP 400 Bad Request if the operation is unknown or parsing fails.
*/
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func calculatehandler(w http.ResponseWriter, r *http.Request) {
	op := r.URL.Query().Get("op")
	astr := r.URL.Query().Get("a")
	bstr := r.URL.Query().Get("b")

	a, err := strconv.Atoi(astr)
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	b, err := strconv.Atoi(bstr)
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	var result int
	switch op {
	case "add":
		result = a + b
	case "subtract":
		result = a - b
	case "multiply":
		result = a * b
	default:
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "Result: ", result)
}
