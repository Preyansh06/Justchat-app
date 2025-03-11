// backend/main.go
package main

import (
	"backend/db"
	"backend/routes"
	"log"
	"net/http"
)

func main() {
	db.InitDB() // Initialize database
	r := routes.SetupRouter()
	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
