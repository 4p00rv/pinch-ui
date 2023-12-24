package main

import (
	"fmt"
	"net/http"

	"github.com/4p00rv/pinch-ui/example/basic"
)

func main() {
	// Define a handler function for root path "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! You've reached the root path.")
	})

	// Define a handler function for path "/hello"
	http.HandleFunc("/basic", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, basic.BasicPage())
	})

	// Start the server, listening on port 8080
	fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err) // Handle errors appropriately in production
	}
}
