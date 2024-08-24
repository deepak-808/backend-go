package controller

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get environment variables
	connectionUrl := os.Getenv("MONGOURI")
	dbName := os.Getenv("DBNAME")
	collectionName := os.Getenv("COLLECTIONNAME")

	fmt.Println("url", connectionUrl)
	clientOption := options.Client().ApplyURI(connectionUrl)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb Connection Successfully Established")
	collection = client.Database(dbName).Collection(collectionName)
}

func GetCollection() *mongo.Collection {
	return collection
}
