// main.go
package main

import (
	"log"
	"net/http"

	"chat-application-backend-go/config"
	"chat-application-backend-go/routes"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()

	r := mux.NewRouter()
	routes.AuthRoutes(r)
	routes.ChatRoutes(r)

	log.Println("Server running on port 5000")
	http.ListenAndServe(":5000", r)
}
