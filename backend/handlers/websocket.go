// backend/handlers/websocket.go
package handlers

import (
	"log"
	"net/http"
	"sync"

	"backend/models"

	"github.com/gorilla/websocket"
)

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Connection manager
var clients = make(map[*websocket.Conn]bool) // Store connected clients
var broadcast = make(chan models.Message)    // Channel for broadcasting messages
var mutex = sync.Mutex{}                     // To prevent race conditions

// HandleWebSocket manages WebSocket connections
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer func() {
		mutex.Lock()
		delete(clients, conn)
		mutex.Unlock()
		conn.Close()
		log.Println("Client disconnected")
	}()

	// Register the client
	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()
	log.Println("New client connected")

	// Listen for messages
	for {
		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Invalid WebSocket message received:", err)
			continue // Keep listening instead of closing the connection
		}
		// if err != nil {
		// 	log.Println("Read error:", err)
		// 	mutex.Lock()
		// 	delete(clients, conn) // Remove disconnected client
		// 	mutex.Unlock()
		// 	break
		// }

		log.Println("Received message:", msg)

		// Store message in database
		err = models.SendMessage(msg.ChatID, msg.SenderID, msg.Content)
		if err != nil {
			log.Println("Database error:", err)
			continue
		}

		// Broadcast message to all clients
		broadcast <- msg
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
