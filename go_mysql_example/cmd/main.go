package main

import (
	"fmt"

	"githhub.com/developerwritings/go_development/gin_example/config"
	"githhub.com/developerwritings/go_development/gin_example/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/user/:id", handlers.GetUserByID)
	router.POST("/user/create", handlers.CreateUser)
	serverConfig := config.Config{}
	serverConfig.NewConfig()
	config := serverConfig.ServerConfig()
	fmt.Println(serverConfig.Port)
	router.Run(":" + config.Port)
}
