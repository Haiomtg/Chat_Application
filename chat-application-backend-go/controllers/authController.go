// controllers/authController.go
package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"chat-application-backend-go/config"
	"chat-application-backend-go/models"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your_jwt_secret")

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hashedPassword)

	collection := config.Client.Database("chat-app").Collection("users")
	collection.InsertOne(context.TODO(), user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("User registered")
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	var foundUser models.User
	collection := config.Client.Database("chat-app").Collection("users")
	err := collection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&foundUser)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)) != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Invalid credentials")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": foundUser.ID,
	})

	tokenString, _ := token.SignedString(jwtKey)
	json.NewEncoder(w).Encode(tokenString)
}
