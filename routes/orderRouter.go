package routes

import (
	controller "github.com/clinton-felix/restaurant-mgt-project/controllers"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/order", controller.GetOrders())
	incomingRoutes.GET("/order/:order_id", controller.GetOrder())
	incomingRoutes.POST("/order", controller.CreateOrder())
	incomingRoutes.PATCH("/order/:order_id", controller.UpdateOrder())
}