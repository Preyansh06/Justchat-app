// backend/handlers/profile.go
package handlers

import (
	"encoding/json"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to your profile!"})
}
