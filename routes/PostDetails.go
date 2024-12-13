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

	var payload model.Details

	c.ShouldBindJSON(&payload)

	fmt.Println("Data Successfully Received from the client")

	var rooms []string

	if payload.Block == "A" {

		collection := client.Database(dbName).Collection("Test")

		rooms = helper.Find(collectionReserve, collection, payload.HourSegment, payload.Block, payload.Day, payload.NumberofHours)

	} else if payload.Block == "B" {

		collection := client.Database(dbName).Collection("B_Block")

		rooms = helper.Find(collectionReserve, collection, payload.HourSegment, payload.Block, payload.Day, payload.NumberofHours)

	} else if payload.Block == "C" {

		collection := client.Database(dbName).Collection("C_Block")

		rooms = helper.Find(collectionReserve, collection, payload.HourSegment, payload.Block, payload.Day, payload.NumberofHours)

	} else if payload.Block == "D" {

		collection := client.Database(dbName).Collection("D_Block")

		rooms = helper.Find(collectionReserve, collection, payload.HourSegment, payload.Block, payload.Day, payload.NumberofHours)

	} else if payload.Block == "H" {

		collection := client.Database(dbName).Collection("H_Block")

		rooms = helper.Find(collectionReserve, collection, payload.HourSegment, payload.Block, payload.Day, payload.NumberofHours)
	} else if payload.Block == "E" {

		collection := client.Database(dbName).Collection("E_Block")

		rooms = helper.Find(collectionReserve, collection, payload.HourSegment, payload.Block, payload.Day, payload.NumberofHours)

	} else if payload.Block == "All" {

		collection := client.Database(dbName).Collection("All_Block")

		rooms = helper.Find(collectionReserve, collection, payload.HourSegment, payload.Block, payload.Day, payload.NumberofHours)

	}

	var length = len(rooms)
	if len(rooms) > 15 {
		length = 15
	}

	response := map[string]interface{}{
		"number":    length,
		"classroom": rooms,
	}

	c.JSON(http.StatusOK, response)

	fmt.Println("Response Sent!")
}

func HandleReserve(client *mongo.Client, c *gin.Context) {

	var res model.Reserve
	c.ShouldBindJSON(&res)

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

	fmt.Printf("Given Username = %s\n", givenUsername)
	fmt.Printf("Given password = %s\n", givenPassword)

	var details []string = helper.CheckUserPassword(collectionCreds, givenUsername)

	fmt.Printf("Real Password = %s\n", details[1])

	if givenPassword != details[1] {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
	} else {
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

}
func HandleSignup(client *mongo.Client, c *gin.Context) {

}
