// backend/main.go
package main

import (
	"backend/db"
	"backend/handlers"
	"backend/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDB() // Initialize database
	r := routes.SetupRouter()
	// Apply JWT Middleware to all routes
	// r.Use(middleware.JWTAuthMiddleware)
	// Print all registered routes
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err == nil {
			log.Println("Registered route:", path)
		}
		return nil
	})
	// Start WebSocket broadcast system
	go handlers.StartBroadcast()

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
