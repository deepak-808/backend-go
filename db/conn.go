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

type MongoDBCollections struct {
	Users    *mongo.Collection
	Orders   *mongo.Collection
	Products *mongo.Collection
}

var client *mongo.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize MongoDB client
	connectionUrl := os.Getenv("MONGOURI")
	clientOptions := options.Client().ApplyURI(connectionUrl)
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("MongoDB Connection Successfully Established")
}

func GetCollection() *MongoDBCollections {
	var dbName = os.Getenv("DBNAME")
	return &MongoDBCollections{
		Users:    client.Database(dbName).Collection("users"),
		Orders:   client.Database(dbName).Collection("orders"),
		Products: client.Database(dbName).Collection("products"),
	}
}
