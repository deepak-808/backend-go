package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionUrl = "mongodb://localhost:27017"
const dbName = "testingdb"
const collectionName = "testingcollection"

var collection *mongo.Collection

func init() {
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
