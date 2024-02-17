package handler

import (
	"errors"
	"fmt"
	"net/http"
	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func PostBookHandler(context *gin.Context) {
	var book book.BookInput

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

func QueryHandler(context *gin.Context) {
	category := context.Query("category")
	id := context.Query("id")

	context.JSON(http.StatusOK, gin.H{
		"id":       id,
		"category": category,
	})
}

func ParamHandler(context *gin.Context) {
	category := context.Param("category")
	id := context.Param("id")

	context.JSON(http.StatusOK, gin.H{
		"id":       id,
		"category": category,
	})
}

func RootHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"title": "Hello World!",
		"bio":   "Lets learn Golang",
	})
}
