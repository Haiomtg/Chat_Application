// models/message.go
package models

import "time"

type Message struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	User      string    `json:"user" bson:"user"`
	Text      string    `json:"text" bson:"text"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}
