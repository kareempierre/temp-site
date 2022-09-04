package main

import (
	"fmt"

	"github.com/alloy-test/alloy-service"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Running ls")
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := alloy.Pool.Ping()
	if err != nil {
		fmt.Println("Error pinging db")
	}

	r.Run(":8888")
}
