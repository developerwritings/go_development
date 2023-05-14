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
	fmt.Println(id)

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User Not Found"})
}
