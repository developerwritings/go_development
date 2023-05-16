package main

import (
	"fmt"

	"githhub.com/developerwritings/go_development/go_mysql_example/config"
	"githhub.com/developerwritings/go_development/go_mysql_example/db"
	"githhub.com/developerwritings/go_development/go_mysql_example/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/user/:id", handlers.GetUserByID)
	router.POST("/user/create", handlers.CreateUser)
	serverConfig := config.Config{}
	serverConfig.NewConfig()
	config := serverConfig.ServerConfig()
	_, err := db.DB()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(serverConfig.Port)
	router.Run(":" + config.Port)
}
