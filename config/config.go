package config

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err!= nil {
        log.Fatal(err)
    }

	err = client.Ping(context.Background(), nil)
	if err!= nil {
        log.Fatal(err)
    }
	log.Println("Connected to MongoDB!")
	DB = client.Database("FirstProyectWebEngineering")
	return DB
}