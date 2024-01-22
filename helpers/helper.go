package helpers

import (
	"context"
	"fmt"
	"log"
	"slices"
	"strconv"

	"github.com/dev.maan707/golang/tests/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Find(collection_forReserve, collection *mongo.Collection, hour int, Block string, Day int) []string {

	fmt.Printf("HourSegment = %d\nBlock = %s\nDay= %d\n", hour, Block, Day)

	if err := collection.Database().Client().Ping(context.Background(), nil); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	//Describing the filter

	filter := bson.M{
		strconv.Itoa(hour): bson.M{"$regex": "(TRAINING|LAB)$"},
		"Day_Key":          Day,
	}

	//Initiating the Find Operation
	fmt.Println("Initiating Filter")
	cursor, err := collection.Find(context.Background(), filter)

	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("Got cursor , Searching values: ")

	defer cursor.Close(context.Background())

	//Checking the length of the cursor
	fmt.Println("Cursor Count:", cursor.RemainingBatchLength())

	//Defining slices
	var ResRooms []string
	var ResHour []int

	if collection_forReserve != nil {
		//Getting all Reserved Data
		cursorReserved, err := collection_forReserve.Find(context.Background(), bson.D{})
		if err != nil {
			log.Fatal(err)
		}
		var ResPayload models.Reserve

		for cursorReserved.Next(context.Background()) {
			fmt.Println("Decoding the Reserved Data")

			err := cursorReserved.Decode(&ResPayload)

			if err != nil {
				log.Fatal("Error occurred while Decoding")
			}

			fmt.Printf("Reserved Room = %s Hour = %d\n", ResPayload.Room_No, ResPayload.Hour)
			ResRooms = append(ResRooms, ResPayload.Room_No)
			ResHour = append(ResHour, ResPayload.Hour)
		}

	}

	//Iterating through the results
	var rooms []string
	var data models.Received

	//Trying to iterate through the received data
	for cursor.Next(context.Background()) {

		fmt.Println("Decoding the json")
		err := cursor.Decode(&data)

		if err != nil {
			log.Fatal(err)
		}

		//printing the object ID just incase
		fmt.Println(data.ID)

		if slices.Contains(ResHour, hour) {
			if slices.Contains(ResRooms, data.RoomNo) {
				continue
			} else {
				rooms = append(rooms, data.RoomNo)
			}
		} else {
			rooms = append(rooms, data.RoomNo)
		}

	}

	//Iterating through the rooms just incase
	for _, room := range rooms {
		fmt.Println(room)
	}
	//Returning the slice back to routes/PostDetails.go
	return rooms

}

func UpdateReserve(collection *mongo.Collection, Hour int, Room_No string) (success string, err error) {

	if err := collection.Database().Client().Ping(context.Background(), nil); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	fmt.Println("Trying to add data into Database")

	ctx := context.TODO()

	docs := models.Reserve{Room_No: Room_No, Hour: Hour}

	fmt.Println("Initiating The cursor")
	cursor, err := collection.InsertOne(ctx, docs)

	if err != nil {
		log.Fatal("Failed in getting cursor", err)
	}

	fmt.Printf("Successfully Added with _id = %s\n", cursor.InsertedID)

	return "y", nil
}

func CheckUserPassword(collection *mongo.Collection, givenUsername string) []string {

	fmt.Printf("Data :\nUsername = %s\n", givenUsername)

	fmt.Println("Applying Filter")
	filter := bson.M{
		"Email": givenUsername,
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Cursor Count:", cursor.RemainingBatchLength())

	defer cursor.Close(context.Background())

	if cursor.RemainingBatchLength() == 0 {
		return []string{"false", "false"}
	}
	var Received models.ReceivedCredentials
	cursor.Decode(&Received)
	fmt.Println("Sending data back")
	fmt.Printf("Password Received = %s\n", Received.Password)
	return []string{Received.Name, Received.Password}

}
