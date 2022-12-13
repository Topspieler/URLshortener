package main

import (
	"URLshortener/configs"
	"URLshortener/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// set new default router with recovery and logger
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Hello from Gin-gonic & mongoDB",
		})
	})
	configs.ConnectDB()

	routes.OrderRoute(router)

	//start and run server
	router.Run(":8080")

}
