package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func InitiateServer() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}

func main() {
	InitiateServer()
	log.Println("Go server")
}
