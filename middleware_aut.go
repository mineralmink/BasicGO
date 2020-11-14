package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	log.Println("in helloHandler")
	c.JSON(http.StatusOK, gin.H{"message": "hello"})
}

func authMiddleware(c *gin.Context) {
	log.Println("start middleware")
	authKey := c.GetHeader("Authorization")
	if authKey != "Bearer token123" {
		c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		c.Abort()
		return
	}

	c.Next()
	log.Println("end middleware")
}

func main() {
	r := gin.Default()
	r.Use(authMiddleware)
	r.GET("/hello", helloHandler)
	r.Run(":1234")
}
