// handler/user.go
package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/yourusername/ecommerce-app/models"
)

type UserHandler struct {
	DB *sql.DB
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{DB: db}
}

func (uh *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	// Your logic to fetch all users from the database
}

func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Your logic to fetch a single user by ID from the database
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Your logic to create a new user in the database
}

func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Your logic to update an existing user in the database
}

func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Your logic to delete an existing user from the database
}
