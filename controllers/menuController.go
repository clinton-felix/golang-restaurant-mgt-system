package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/clinton-felix/restaurant-mgt-project/database"
	"github.com/clinton-felix/restaurant-mgt-project/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// create menucollection in database
var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

// get all the menu which exists in a database
func GetMenus() gin.HandlerFunc{
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		result, err := menuCollection.Find(context.TODO(), bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured when listing menu items"})
		}

		var allMenus []bson.M
		if err = result.All(ctx, &allMenus); err != nil {
			c.JSON(http.StatusOK, allMenus)
		}
	}
}

func GetMenu() gin.HandlerFunc{
	return func(c *gin.Context) {
		// set a timeout duration for the database connection
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
		menuId := c.Param("menu_id")
		var menu models.Menu

		err := foodCollection.FindOne(ctx, bson.M{"menu_id": menuId}).Decode(&menu)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while fetching the menu"})
		}
		c.JSON(http.StatusOK, menu)
	}
}

func CreateMenu() gin.HandlerFunc{
	return func(c *gin.Context) {
		var menu models.Menu
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
		// bind the menu JSOn 
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H {"error":err.Error()})
			return
		}

		// validate the menu model
		validationErr := validate.Struct(menu)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H {"error":validationErr.Error()})
			return
		}

		menu.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.ID = primitive.NewObjectID()
		menu.Menu_id = menu.ID.Hex()
		
		result, insertErr := menuCollection.InsertOne(ctx, menu)
		if insertErr != nil {
			msg := fmt.Sprintf("menu item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
		defer cancel()
	}
}


// update menu
func UpdateMenu() gin.HandlerFunc{
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu

		// bind to the Menu Object
		if err := c.BindJSON(&menu) ; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		menuId := c.Param("menu_id")
		filter := bson.M{"menu_id":menuId}

		// updateObj is a varable that effects update on the
		// mongo database
		var updateObj primitive.D

		// confirm that menu object has a start and end date
		if menu.Start_date != nil && menu.End_date != nil {
			if !inTimeSpan(*menu.Start_date, *menu.End_date, time.Now()){
				msg := "kindly retype the time"
				c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
				defer cancel()
				return
			}

			// append to the update object
			updateObj = append(updateObj, bson.E{"start_date", menu.Start_date})
			updateObj = append(updateObj, bson.E{"end_date", menu.End_date})

			if menu.Name != "" {
				updateObj = append(updateObj, bson.E{"name", menu.Name})
			}
			if menu.Category != "" {
				updateObj = append(updateObj, bson.E{"name", menu.Category})
			}

			menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			updateObj = append(updateObj, bson.E{"updated_at", menu.Updated_at})

			upsert := true

			opt := options.UpdateOptions{
				Upsert: &upsert,
			}

			// attempt to update the database with updateObj
			result, err := menuCollection.UpdateOne(
				ctx,
				filter,
				bson.D{
					{"$set", updateObj},
				},
				&opt,
			)

			if err != nil{
				msg := "menu update failed"
				c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
			}

			defer cancel()
			c.JSON(http.StatusOK, result)
		}

	}
}