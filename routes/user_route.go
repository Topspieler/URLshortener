package routes

import (
	"URLshortener/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoute(router *gin.Engine) {
	router.POST("/api/order", controllers.CreateOrder())
	router.GET("/api/order/:orderId", controllers.GetAOrder())
	router.DELETE("/api/order/:orderId", controllers.DeleteAOrder())
	router.GET("/api/orders", controllers.GetAllOrders())

}
