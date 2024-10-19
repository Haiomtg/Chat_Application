package main

import (
	"fmt"
	"net/http"

	"github.com/Haiomtg/chatapp/router"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Set up routes
	router.SetupRoutes(r)

	fmt.Println("Chat server started on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
