// backend/handlers/auth.go
package handlers

import (
	"encoding/json"
	"net/http"

	"backend/models"
	"backend/utils"
)

// Signup handles user registration
func Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	if err := models.CreateUser(user); err != nil {
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	if err := models.CreateUser(user); err != nil {
		http.Error(w, "Could not create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

// Login handles user authentication
func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Fetch user from database (for now, placeholder logic)
	storedUser, err := models.GetUserByEmail(user.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Verify password
	if !utils.CheckPasswordHash(user.Password, storedUser.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(storedUser.Email)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	// Respond with the token
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
