package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/golang-jwt/jwt" // Import updated JWT package
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var clients = make(map[*websocket.Conn]bool)
var mu sync.Mutex

// Secret key for JWT
var secretKey = []byte("your_secret_key")

// Function to generate a token
func generateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})
	return token.SignedString(secretKey)
}

// Middleware to check token
func checkToken(r *http.Request) (string, error) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return "", fmt.Errorf("no token provided")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	return token.Claims.(jwt.MapClaims)["username"].(string), nil
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	username, err := checkToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error while upgrading connection:", err)
		return
	}
	defer conn.Close()

	mu.Lock()
	clients[conn] = true
	// Replace fmt.Fprintf with websocket.Message.Send
	err = conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Welcome %s!", username))) // Use the username here
	if err != nil {
		// Handle the error appropriately
		fmt.Println("Error sending message:", err)
	}
	mu.Unlock()
}

func broadcastMessage(msg []byte) {
	mu.Lock()
	defer mu.Unlock()
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			fmt.Println("Error while writing message:", err)
			client.Close()
			delete(clients, client)
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnection)
	fmt.Println("Chat server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
