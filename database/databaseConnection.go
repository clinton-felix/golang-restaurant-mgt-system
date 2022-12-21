package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	MongoDb := "mongodb://localhost:27017"
	fmt.Print(MongoDb)

	// create a new mongo client; also catch error
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// pass contect to connect to the database
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("connected to mongo DB")
	return client
}

// creating DB instance by calling the function DBinstance(), and capture
// value in a variable called Client
var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)

	return collection
}