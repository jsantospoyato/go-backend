package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Creates a basic /GET endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Print Hello World
		fmt.Fprint(w, "Hello world")
	})

	fmt.Println("🚀 Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
