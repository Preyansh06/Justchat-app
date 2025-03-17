// backend/models/chat_participant.go
package models

import (
	"backend/db"
	"log"
)

// AddUserToChat adds a user to a chat
func AddUserToChat(chatID, userID int) error {
	query := "INSERT INTO chat_participants (chat_id, user_id) VALUES ($1, $2) ON CONFLICT DO NOTHING"
	_, err := db.DB.Exec(query, chatID, userID)
	if err != nil {
		log.Println("Error adding user to chat:", err)
		return err
	}
	return nil
}
