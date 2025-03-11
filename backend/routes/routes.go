// backend/routes/routes.go
package routes

import (
	"backend/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/signup", handlers.Signup).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	// r.HandleFunc("/login", handlers.Login).Methods("POST")
	return r
}
