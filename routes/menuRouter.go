package routes

import (
	controller "github.com/clinton-felix/restaurant-mgt-project/controllers"
	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menu", controller.GetMenus())
	incomingRoutes.GET("/menu/:menu_id", controller.GetMenu())
	incomingRoutes.POST("/menu", controller.CreateMenu())
	incomingRoutes.PATCH("/menu/:menu_id", controller.UpdateMenu())
}