package main

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	println("Shereats Main")
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		InsertDatabase()
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.Run()
}
