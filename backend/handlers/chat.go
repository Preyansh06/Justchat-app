// backend/handlers/chat.go
package handlers

import (
	"backend/models"
	"encoding/json"
	"net/http"
)

// CreateChatHandler creates a private or group chat
func CreateChatHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		IsGroup bool  `json:"is_group"`
		UserIDs []int `json:"user_ids"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	chatID, err := models.CreateChat(req.IsGroup)
	if err != nil {
		http.Error(w, "Could not create chat", http.StatusInternalServerError)
		return
	}

	// Add users to chat
	for _, userID := range req.UserIDs {
		models.AddUserToChat(chatID, userID)
	}

	json.NewEncoder(w).Encode(map[string]int{"chat_id": chatID})
}
