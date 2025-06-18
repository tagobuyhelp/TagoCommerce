package database

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    
    "TagoCommerce/internal/config"
)

var Client *mongo.Client
var Database *mongo.Database

// ConnectDB establishes a connection to the MongoDB database
func ConnectDB() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    clientOptions := options.Client().ApplyURI(config.AppConfig.MongoURI)
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal("Failed to connect to MongoDB:", err)
    }

    // Check the connection
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal("Failed to ping MongoDB:", err)
    }

    log.Println("Connected to MongoDB!")
    Client = client
    Database = client.Database("tagocommerce")
}

// GetCollection returns a handle to the specified collection
func GetCollection(collectionName string) *mongo.Collection {
    return Database.Collection(collectionName)
}