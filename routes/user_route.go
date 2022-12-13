package routes

import (
	"URLshortener/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoute(router *gin.Engine) {
	router.POST("/order", controllers.CreateOrder())
	router.GET("/order/:orderId", controllers.GetAOrder())
	router.DELETE("/order/:orderId", controllers.DeleteAOrder())
	router.GET("/orders", controllers.GetAllOrders())

}
