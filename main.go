package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	// :nameから値を取り出す
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("hello %s", name),
	})
}

func instruments(c *gin.Context) {
	var instruments [3]string
	instruments[0] = "piano"
	instruments[1] = "guitar"
	instruments[2] = "bass"

	c.JSON(200, gin.H{
		"message": instruments,
	})
}

func games(c *gin.Context) {
	var games [2]string
	games[0] = "BanG Dream!"
	games[1] = "ragnarokorigin"

	c.JSON(200, gin.H{
		"message": games,
	})
}

func main() {
	router := gin.Default()

	router.GET("/:name", home)
	router.GET("/instruments", instruments)
	router.GET("/games", games)

	router.Run(":5000")
}
