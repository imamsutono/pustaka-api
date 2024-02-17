package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	Title string      `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
}

func postBookHandler(context *gin.Context) {
	var book Book

	err := context.ShouldBindJSON(&book)
	if err != nil {
		var ve *validator.ValidationErrors
		errorMessages := []string{}
		if errors.As(err, &ve) {
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
		}

		context.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"title": book.Title,
		"price": book.Price,
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
