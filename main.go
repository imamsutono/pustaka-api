package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/books/:category/:id", paramHandler)
	router.GET("/product", queryHandler)

	router.POST("/books", postBookHandler)

	router.Run()
}

type Book struct {
	Title    string
	SubTitle string `json:"sub_title"`
	Price    int
}

func postBookHandler(context *gin.Context) {
	var book Book

	err := context.ShouldBindJSON(&book)
	if err != nil {
		log.Fatal(err)
	}

	context.JSON(http.StatusOK, gin.H{
		"title":     book.Title,
		"sub_title": book.SubTitle,
		"price":     book.Price,
	})
}

func queryHandler(context *gin.Context) {
	category := context.Query("category")
	id := context.Query("id")

	context.JSON(http.StatusOK, gin.H{
		"id":       id,
		"category": category,
	})
}

func paramHandler(context *gin.Context) {
	category := context.Param("category")
	id := context.Param("id")

	context.JSON(http.StatusOK, gin.H{
		"id":       id,
		"category": category,
	})
}

func rootHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"title": "Hello World!",
		"bio":   "Lets learn Golang",
	})
}
