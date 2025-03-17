// backend/handlers/websocket.go
package handlers

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Connection manager
var clients = make(map[*websocket.Conn]bool) // Store connected clients
var broadcast = make(chan Message)           // Channel for broadcasting messages
var mutex = sync.Mutex{}                     // To prevent race conditions

// Message struct for chat messages
type Message struct {
	ChatID   int    `json:"chat_id"`
	SenderID int    `json:"sender_id"`
	Content  string `json:"content"`
}

// HandleWebSocket manages WebSocket connections
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	// Register the client
	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()
	log.Println("New client connected")

	// Listen for messages
	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Read error:", err)
			mutex.Lock()
			delete(clients, conn) // Remove disconnected client
			mutex.Unlock()
			break
		}

		log.Println("Received message:", msg)
		broadcast <- msg // Send message to broadcast channel
	}
}

// StartBroadcast listens for messages and sends them to all connected clients
func StartBroadcast() {
	for {
		msg := <-broadcast
		log.Println("Broadcasting message:", msg)

		mutex.Lock()
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Println("Write error:", err)
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}
