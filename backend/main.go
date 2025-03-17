// backend/main.go
package main

import (
	"backend/db"
<<<<<<< HEAD
=======
	"backend/handlers"
>>>>>>> working-branch
	"backend/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDB() // Initialize database
	r := routes.SetupRouter()
<<<<<<< HEAD
=======
	// Apply JWT Middleware to all routes
	// r.Use(middleware.JWTAuthMiddleware)
>>>>>>> working-branch
	// Print all registered routes
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err == nil {
			log.Println("Registered route:", path)
		}
		return nil
	})
<<<<<<< HEAD
=======
	// Start WebSocket broadcast system
	go handlers.StartBroadcast()
>>>>>>> working-branch

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
