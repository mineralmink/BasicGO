package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	log.Println("in helloHandler")
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}

func middleware(c *gin.Context) {
	log.Println("start middleware2")
	c.Next()
	log.Println("end middleware2")
}

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		log.Println("start middleware")
		c.Next() //helloHandler(c) same
		log.Println("end middleware")
	})
	r.Use(middleware)

	r.GET("/hello", helloHandler)
	r.Run(":1234")
}
