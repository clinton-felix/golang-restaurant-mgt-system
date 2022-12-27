package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/clinton-felix/restaurant-mgt-project/database"
	"github.com/clinton-felix/restaurant-mgt-project/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// foodCollection is a variable of type mongo.Collection
// which stores the value of a database instance of collection "food"
var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var validate = validator.New()


// returns the list of foods in the database
func GetFoods() gin.HandlerFunc{
	return func(c *gin.Context) {

	}
}

func GetFood() gin.HandlerFunc{
	return func(c *gin.Context) {
		// set a timeout duration for the database connection
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
		foodId := c.Param("food_id")
		var food models.Food

		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while fetching the food item from the database"})
		}
		c.JSON(http.StatusOK, food)
	}
}

func CreateFood() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu
		var food models.Food

		// Bind the food model to the header
		if err := c.BindHeader(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
			return
		}

		// validate the food model before proceeding to
		// create the different fields of the food struct model
		validationError := validate.Struct(food)
		if validationError != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationError.Error()})
			return
		}

		// Check first to check (using "FindOne") that the menu item exists
		// using the menu ID before creating the food item
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
		defer cancel()
		if err != nil {
			msg := fmt.Sprintf("menu was not found")
			c.JSON(http.StatusBadRequest, gin.H{"error":msg})
			return
		}
		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.ID = primitive.NewObjectID()
		food.Food_id = food.Food_id.Hex()
		var num = toFixed(*food.Price, 2)
		food.Price = &num

		result, insertErr := foodCollection.InsertOne(ctx, food)
		if insertErr != nil {
			msg := fmt.Sprintf("food item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func UpdateFood() gin.HandlerFunc{
	return func(c *gin.Context) {

	}
}

// function to round up a number
func round(num float64) int {

}

// function to convert to fixed
func toFixed(num float64, precision int) float64{

}