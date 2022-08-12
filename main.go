package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"title": "Hello World!",
			"bio":   "Lets learn Golang",
		})
	})

	router.Run()
}
