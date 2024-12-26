package main

import (
	"github.com/gin-gonic/gin"
	"github.com/victorlazzaroli/microservicesTest/message/api/message"
	"github.com/victorlazzaroli/microservicesTest/message/startup"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	startup.ConnectDatabase()

	// Define a route for the root URL
	router.GET("/", func(c *gin.Context) {
		panic("PAnico")
	})

	message.NewMessageController(router, startup.DB)

	// Run the server on port 8080
	router.Run(":8080")
}
