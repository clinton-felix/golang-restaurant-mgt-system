package routes

import (
	controller "github.com/clinton-felix/restaurant-mgt-project/controllers"
	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controller.GetFoods())
	incomingRoutes.GET("/foods/:food_id", controller.GetFood())
	incomingRoutes.POST("/foods", controller.CreateFood())
	incomingRoutes.PATCH("/food/:food_id", controller.UpdateFood())	// update a food item
}