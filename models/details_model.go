package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Struct defining the data which will be received by the backend from app
type Details struct {
	Block       string `json:"block"`
	RoonType    string `json:"classroom"`
	Day         int    `json:"day"`
	HourSegment int    `json:"hours"`
}

type ColumnsData struct {
	Columns map[string]string
}

//Struct defining the data which will be received by the backend from mongoDB

type Received struct {
	ID      primitive.ObjectID `bson:"_id"`
	RoomNo  string             `bson:"Room_no"`
	DayKey  int                `bson:"Day_key"`
	DayTime string             `bson:"Day/Time"`
	Columns ColumnsData        `bson:",inline"`
}

type Reserve struct {
	Room_No string `bson:"Room_No"`
	Hour    int    `bson:"Hour"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Signup struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type ReceivedCredentials struct {
	ID       primitive.ObjectID `bson:"_id"`
	Serial   int                `bson:"S_No"`
	Name     string             `bson:"Name"`
	Email    string             `bson:"Email"`
	Password string             `bson:"Password"`
}
