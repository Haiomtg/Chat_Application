// routes/authRoutes.go
package routes

import (
	"github.com/gorilla/mux"

	"chat-application-backend-go/controllers"
)

func AuthRoutes(r *mux.Router) {
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
}
