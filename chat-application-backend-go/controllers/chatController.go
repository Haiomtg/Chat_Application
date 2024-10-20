// controllers/chatController.go
package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"chat-application-backend-go/config"
	"chat-application-backend-go/models"

	_ "github.com/go-sql-driver/mysql"
)

func SendMessage(w http.ResponseWriter, r *http.Request) {
	var message models.Message
	json.NewDecoder(r.Body).Decode(&message)
	message.Timestamp = time.Now()

	db := config.DB // Assuming you have a DB field in your config for MySQL
	_, err := db.Exec("INSERT INTO messages (text, timestamp) VALUES (?, ?)", message.Text, message.Timestamp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(message)
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	db := config.DB // Assuming you have a DB field in your config for MySQL
	rows, err := db.Query("SELECT text, timestamp FROM messages")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var message models.Message
		if err := rows.Scan(&message.Text, &message.Timestamp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		messages = append(messages, message)
	}

	json.NewEncoder(w).Encode(messages)
}
