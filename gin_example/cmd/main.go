package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type user struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"artist"`
}

func main() {
	router := gin.Default()
	router.GET("/", Homepage)
	router.GET("/user/:id", getUserByID)
	router.POST("/user/create", CreateUser)

	router.Run("localhost:8080")
}

func Homepage(c *gin.Context) {
}

func CreateUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

}

func getUserByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User Not Found"})
}
