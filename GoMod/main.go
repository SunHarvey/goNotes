package main

import (
	"github.com/gin-gonic/gin"
    "Gone/testmod"
)

func main() {
    testmod.Hello()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
