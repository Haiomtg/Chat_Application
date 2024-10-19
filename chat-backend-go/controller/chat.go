package controller

import (
	"encoding/json"
	"net/http"

	"github.com/yourusername/chat-backend-go/middleware"
	"github.com/yourusername/chat-backend-go/model"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if model.AuthenticateUser(user.Username, user.Password) {
		token, err := middleware.GenerateToken(user.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(token))
	} else {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := model.RegisterUser(user.Username, user.Password); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
