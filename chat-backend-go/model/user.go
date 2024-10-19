package model

import (
	"errors"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = make(map[string]string) // Simulating a user database

func RegisterUser(username, password string) error {
	if _, exists := users[username]; exists {
		return errors.New("user already exists")
	}
	users[username] = password
	return nil
}

func AuthenticateUser(username, password string) bool {
	if pass, exists := users[username]; exists {
		return pass == password
	}
	return false
}
