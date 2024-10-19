package router

import (
	"github.com/gorilla/mux"
	"github.com/yourusername/chat-backend-go/controller"
	"github.com/yourusername/chat-backend-go/middleware"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/ws", middleware.CheckToken(controller.HandleConnection)).Methods("GET")
	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/register", controller.Register).Methods("POST")
}
