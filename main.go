package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", pinghandler)
	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/count", counthandler)
	http.HandleFunc("/calculate", calculatehandler)
	http.HandleFunc("/agent", agenthandler)
	http.HandleFunc("/dashboard", dashboardhandler)

	http.HandleFunc("/legacy", legacyhandler)
	http.HandleFunc("/v2", v2handler)

	fmt.Println("server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
