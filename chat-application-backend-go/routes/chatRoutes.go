// routes/chatRoutes.go
package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"chat-application-backend-go/controllers"
	"chat-application-backend-go/middleware"
)

func ChatRoutes(r *mux.Router) {
	r.Handle("/messages", middleware.AuthMiddleware(http.HandlerFunc(controllers.SendMessage))).Methods("POST")
	r.Handle("/messages", middleware.AuthMiddleware(http.HandlerFunc(controllers.GetMessages))).Methods("GET")
}
