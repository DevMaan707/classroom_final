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

	//client options
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)

	//Above two lines are directly copied from the documentation provided by mongoDB

	//connect to mongoDB
	client, err := mongo.Connect(context.TODO(), opts)

	//Handling Errors which may raise
	if err != nil {
		log.Fatal((err))
		return nil, err
	}

	fmt.Println("Successful Connection with MongoDB")

	return client, nil
}
