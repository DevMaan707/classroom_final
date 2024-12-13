package controllers

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://AashishReddy:test123@cluster0.wd6ydng.mongodb.net/?retryWrites=true&w=majority"

func ConnectDB() (*mongo.Client, error) {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal((err))
		return nil, err
	}

	fmt.Println("Successful Connection with MongoDB")

	return client, nil
}
