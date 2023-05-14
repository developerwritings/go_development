package handlers

import (
	"fmt"
	"net/http"

	"githhub.com/developerwritings/go_development/gin_example/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	err := db.QueryRow("SELECT id, name FROM tags where id = ?", id).Scan(&User.ID, &User.Name)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User Not Found"})
}
