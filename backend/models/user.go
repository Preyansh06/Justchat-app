// backend/models/user.go
package models

import (
	"backend/db"
	"log"
)

// User struct represents a user in the database
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"` // Hashed password
}

// CreateUser inserts a new user into the database
func CreateUser(user User) error {
	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)"
	_, err := db.DB.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		log.Println("Error inserting user:", err)
		return err
	}
	return nil
}

// GetUserByEmail fetches a user from the database by email
func GetUserByEmail(email string) (User, error) {
	var user User
	query := "SELECT id, username, email, password FROM users WHERE email = $1"
	err := db.DB.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	return user, err
}
