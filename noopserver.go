package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Noop server responding",
			"ts":      time.Now()})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
