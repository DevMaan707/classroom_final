package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"

	helper "github.com/dev.maan707/golang/tests/helpers"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"

	model "github.com/dev.maan707/golang/tests/models"
	"github.com/gin-gonic/gin"
)

const dbName = "TimeTable"

// const colName = "A_Block"
const reserve = "Reserve"

const LoginCol = "Login_Credentials"

var collectionReserve *mongo.Collection

func HandleData(c *gin.Context, client *mongo.Client) {

	//creating instance of the structure
	var payload model.Details

	//binding the structure to the received data
	c.ShouldBindJSON(&payload)

	//Printing received
	fmt.Println("Data Successfully Received from the client")

	//Defining "rooms" var

	var rooms []string

	//Conditional Programming to redirect to different blocks

	if payload.Block == "A" {

		//Getting Collection
		collection := client.Database(dbName).Collection("A_Block")

		//Getting the search data from mongoDB in "rooms"

		rooms = helper.Find(collectionReserve, collection, payload.HourSegment, payload.Block, payload.Day)

	} else if payload.Block == "B" {

		//Getting Collection

		collection := client.Database(dbName).Collection("B_Block")

		//Getting the search data from mongoDB in "rooms"

		rooms = helper.Find(collectionReserve, collection, payload.HourSegment, payload.Block, payload.Day)

	} else if payload.Block == "C" {

		//Getting COllection

		collection := client.Database(dbName).Collection("C_Block")

		//Getting the search data from mongoDB in "rooms"

		rooms = helper.Find(collectionReserve, collection, payload.HourSegment, payload.Block, payload.Day)

	} else if payload.Block == "D" {

		//Getting Collection

		collection := client.Database(dbName).Collection("D_Block")

		//Getting the search data from mongoDB in "rooms"

		rooms = helper.Find(collectionReserve, collection, payload.HourSegment, payload.Block, payload.Day)

	} else if payload.Block == "H" {

		//Getting COllection

		collection := client.Database(dbName).Collection("H_Block")

		//Getting the search data from mongoDB in "rooms"

		rooms = helper.Find(collectionReserve, collection, payload.HourSegment, payload.Block, payload.Day)
	} else if payload.Block == "E" {

		//Getting Collection

		collection := client.Database(dbName).Collection("E_Block")

		//Getting the search data from mongoDB in "rooms"

		rooms = helper.Find(collectionReserve, collection, payload.HourSegment, payload.Block, payload.Day)

	} else if payload.Block == "All" {

		//Getting Collection
		collection := client.Database(dbName).Collection("All_Block")

		//Getting the search data from mongoDB in "rooms"

		rooms = helper.Find(collectionReserve, collection, payload.HourSegment, payload.Block, payload.Day)

	}

	//Limiting the search results to only 5 Rooms

	var length = 5
	if len(rooms) < 5 {
		length = len(rooms) - 1
	}

	//Creating interface which consists of list of rooms and the length of the list

	response := map[string]interface{}{
		"number":    length,
		"classroom": rooms,
	}

	//Finally , sending the data back to the application
	c.JSON(http.StatusOK, response)

	fmt.Println("Response Sent!")
}

func HandleReserve(client *mongo.Client, c *gin.Context) {

	//creating the instance of the Reserve struct
	var res model.Reserve

	//Binding the Request with the Reserve Struct
	c.ShouldBindJSON(&res)

	//Getting the Reserve Collection
	collectionReserve = client.Database(dbName).Collection(reserve)

	_, err := helper.UpdateReserve(collectionReserve, res.Hour, res.Room_No)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully Added Class %s for the HourSegment %d\n", res.Room_No, res.Hour)

}
func HandleLogin(client *mongo.Client, c *gin.Context) {

	collectionCreds := client.Database(dbName).Collection(LoginCol)

	var credentials model.Login
	err := c.ShouldBindJSON(&credentials)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	var givenUsername = credentials.Username
	var givenPassword = credentials.Password
	var details []string = helper.CheckUserPassword(collectionCreds, givenUsername)

	if givenPassword != details[1] {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
	}

	expiration := time.Now().Add(time.Hour * 200)

	claims := &model.Claims{
		Username: givenUsername,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}
	var jwt_Key = []byte("secret_key")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwt_Key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})

}
func HandleSignup(client *mongo.Client, c *gin.Context) {

}
