package routes

import (
	controller "github.com/clinton-felix/restaurant-mgt-project/controllers"
	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/table", controller.GetTables())
	incomingRoutes.GET("/table/:table_id", controller.GetTable())
	incomingRoutes.POST("/table", controller.CreateTable())
	incomingRoutes.PATCH("/table/:table_id", controller.UpdateTable())
}