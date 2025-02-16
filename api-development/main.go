package main

import (
	"fmt"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home page")
	fmt.Println("Endpoint hit: homepage")
}

// go doc http
func main() {
	http.HandleFunc("/", homepage)
	http.ListenAndServe("localhost:8093", nil)
}
