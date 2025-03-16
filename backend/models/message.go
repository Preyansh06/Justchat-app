// backend/models/message.go
package models

import (
	"backend/db"
	"log"
)

// Message struct represents a chat message
type Message struct {
	ID       int    `json:"id"`
	ChatID   int    `json:"chat_id"`
	SenderID int    `json:"sender_id"`
	Content  string `json:"content"`
}

// SendMessage stores a message in the database
func SendMessage(chatID, senderID int, content string) error {
	query := "INSERT INTO messages (chat_id, sender_id, content) VALUES ($1, $2, $3)"
	_, err := db.DB.Exec(query, chatID, senderID, content)
	if err != nil {
		log.Println("Error sending message:", err)
		return err
	}
	return nil
}
