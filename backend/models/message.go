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
	SentAt   string `json:"sent_at"`
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

// func SendMessage(chatID, senderID int, content string) error {
// 	if !ChatExists(chatID) {
// 		log.Println("ERROR: Chat ID does not exist:", chatID)
// 		return errors.New("chat does not exist")
// 	}

// 	query := "INSERT INTO messages (chat_id, sender_id, content) VALUES ($1, $2, $3)"
// 	_, err := db.DB.Exec(query, chatID, senderID, content)
// 	if err != nil {
// 		log.Println("Database error:", err)
// 		return err
// 	}
// 	return nil
// }

// GetMessages fetches messages for a given chat ID
func GetMessages(chatID int) ([]Message, error) {
	query := "SELECT id, chat_id, sender_id, content, sent_at FROM messages WHERE chat_id = $1 ORDER BY sent_at ASC"
	rows, err := db.DB.Query(query, chatID)
	if err != nil {
		log.Println("Error fetching messages:", err)
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		if err := rows.Scan(&msg.ID, &msg.ChatID, &msg.SenderID, &msg.Content, &msg.SentAt); err != nil {
			log.Println("Error scanning message row:", err)
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, nil
}

// func ChatExists(chatID int) bool {
// 	var exists bool
// 	query := "SELECT EXISTS(SELECT 1 FROM chats WHERE id=$1)"
// 	db.DB.QueryRow(query, chatID).Scan(&exists)
// 	return exists
// }
