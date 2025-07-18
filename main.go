package main

import (
	"fmt"
	"net/http"

	"go-backend/user/domain"
)

// /POST method
// Uses 'name' to greet the User via QueryParams (r.Url.Query()...)
func greet() {
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		newUser := domain.User{
			Name: r.URL.Query().Get("name"),
		}

		fmt.Fprintf(w, newUser.Greet())
	})
}

func main() {
	// Creates a basic /GET endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Print Hello World
		fmt.Fprint(w, "Hello world")
	})

	greet()

	fmt.Println("🚀 Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
