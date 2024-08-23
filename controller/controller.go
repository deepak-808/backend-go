package controller

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionUrl = os.Getenv("MONGOURI")
var dbName = os.Getenv("DBNAME")
var collectionName = os.Getenv("COLLECTIONNAME")

var collection *mongo.Collection

func init() {
	fmt.Println("url", connectionUrl)
	clientOption := options.Client().ApplyURI(connectionUrl)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb Connection Sucessfully")
	collection = client.Database(dbName).Collection(collectionName)
}
func GetCollection() *mongo.Collection {
	return collection
}
