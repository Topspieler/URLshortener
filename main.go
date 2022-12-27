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
		c.File("./public/index.html")
	})
	configs.ConnectDB()

	routes.OrderRoute(router)

	//start and run server
	router.Run(":8080")

}
