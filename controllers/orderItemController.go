package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrderItems() gin.HandlerFunc{
	return func(c *gin.Context) {

	}
}

func GetOrderItemsByOrder() gin.HandlerFunc{
	return func(c *gin.Context) {

	}
}

// function to return the item by Order
func ItemsByOrder(id string) (OrderItems []primitive.H, err error) {

}


func GetOrderItem() gin.HandlerFunc{
	return func(c *gin.Context) {
		
	}
}

func CreateOrderItem() gin.HandlerFunc{
	return func(c *gin.Context) {
		
	}
}

func UpdateOrderItem() gin.HandlerFunc{
	return func(c *gin.Context) {
		
	}
}