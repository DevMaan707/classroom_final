package main

import (
	"fmt"
	"log"

	DB "github.com/dev.maan707/golang/tests/controllers"
	route "github.com/dev.maan707/golang/tests/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	client, err := DB.ConnectDB()

	if err != nil {
		log.Fatal("Failed To connect to MongoDB: ", err)
	}
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		fmt.Println("Success")
	})

	router.POST("/room-details", func(c *gin.Context) {
		route.HandleData(c, client)
	})

	router.POST("/reserve", func(c *gin.Context) {
		route.HandleReserve(client, c)
	})

	router.POST("/login", func(c *gin.Context) {
		route.HandleLogin(client, c)
	})

	router.POST("/signup", func(c *gin.Context) {
		route.HandleSignup(client, c)
	})
	router.Run(":6919")
}
