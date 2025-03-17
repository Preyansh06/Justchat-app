// backend/model/chat.go
package models

import (
	"backend/db"
	"log"
)

// Chat structure represents if (private or group chat)
type Chat struct {
	ID      int  `json:"id"`
	IsGroup bool `json:"is_group"`
}

// CreateChat creates a new chat and returns the chat ID
func CreateChat(isGroup bool) (int, error) {
	var chatID int
	query := "INSERT INTO chats (is_group) VALUES ($1) RETURNING id"
	err := db.DB.QueryRow(query, isGroup).Scan(&chatID)
	if err != nil {
		log.Println("Error creating chat:", err)
		return 0, err
	}
	return chatID, nil
}
