package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello World! %s", time.Now())
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "Hello World! %s", time.Now())
	case "/about":
		fmt.Fprintf(w, "This is the about page.")
	case "/contact":
		fmt.Fprintf(w, "This is the contact page.")
	case "/home":
		fmt.Fprintf(w, "This is the home page.")
	default:
		fmt.Fprintf(w, "404 Not Found")
	}
	fmt.Printf("Handling function with %s request\n", r.Method)
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
