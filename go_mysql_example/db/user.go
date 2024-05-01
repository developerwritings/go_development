// handler/user_handler.go
package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	DB *sql.DB
}

// NewUserHandler initializes a new UserHandler
func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{DB: db}
}

// GetUsers handles GET request to fetch all users
func (uh *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := uh.DB.Query("SELECT * FROM users")
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []map[string]interface{}
	for rows.Next() {
		user := make(map[string]interface{})
		if err := rows.Scan(&user["id"], &user["username"], &user["email"]); err != nil {
			http.Error(w, "Failed to scan user", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Failed to iterate over users", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, users)
}

// GetUser handles GET request to fetch a single user by ID
func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Your logic to fetch a single user by ID from the database
}

// CreateUser handles POST request to create a new user
func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	_, err = uh.DB.Exec("INSERT INTO users (username, email) VALUES (?, ?)", userData["username"], userData["email"])
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	respondWithMessage(w, http.StatusCreated, "User created successfully")
}

// UpdateUser handles PUT request to update an existing user
func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Your logic to update an existing user in the database
}

// DeleteUser handles DELETE request to delete an existing user
func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Your logic to delete an existing user from the database
}

// Helper function to respond with JSON data
func respondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to encode response data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

// Helper function to respond with a plain text message
func respondWithMessage(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	w.Write([]byte(message))
}
