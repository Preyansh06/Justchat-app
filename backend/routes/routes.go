// backend/routes/routes.go
package routes

import (
	"backend/handlers"
<<<<<<< HEAD
=======
	"backend/middleware"
>>>>>>> working-branch

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	// r := mux.NewRouter()
	r := mux.NewRouter().StrictSlash(true) // Add StrictSlash
	r.HandleFunc("/signup", handlers.Signup).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	// r.HandleFunc("/login", handlers.Login).Methods("POST")
<<<<<<< HEAD
=======

	r.HandleFunc("/ws", handlers.HandleWebSocket)
	// Protected route (requires JWT)
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.JWTAuthMiddleware)
	protected.HandleFunc("/profile", handlers.Profile).Methods("GET")
	protected.HandleFunc("/chat", handlers.CreateChatHandler).Methods("POST")
	protected.HandleFunc("/chat/{id}/message", handlers.SendMessageHandler).Methods("POST")
	protected.HandleFunc("/chat/{id}/messages", handlers.GetMessagesHandler).Methods("GET")
>>>>>>> working-branch
	return r
}
