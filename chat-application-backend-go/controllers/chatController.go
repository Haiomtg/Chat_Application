// controllers/chatController.go
package controllers

import (
    "context"
    "encoding/json"
    "net/http"
    "time"

	"chat-application-backend-go/config"
	"chat-application-backend-go/models"

	"go.mongodb.org/mongo-driver/bson"
)

func SendMessage(w http.ResponseWriter, r *http.Request) {
    var message models.Message
    json.NewDecoder(r.Body).Decode(&message)
    message.Timestamp = time.Now()

    collection := config.Client.Database("chat-app").Collection("messages")
    collection.InsertOne(context.TODO(), message)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(message)
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
    collection := config.Client.Database("chat-app").Collection("messages")
    cursor, _ := collection.Find(context.TODO(), bson.M{})
    var messages []models.Message
    cursor.All(context.TODO(), &messages)

    json.NewEncoder(w).Encode(messages)
}

