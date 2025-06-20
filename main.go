package main

import (
	"Golang-API-Authentication/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariavles()
	initializers.ConnectToDb()
	initializers.SyncDb()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
