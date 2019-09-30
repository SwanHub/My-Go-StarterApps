package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{greeting: Hello World}")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About!</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about) // exact paths!
	fmt.Println("Server Starting...")
	http.ListenAndServe(":3000", nil)
}

// great for REST APIs
