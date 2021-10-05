package handler

import (
	"net/http"

	"pustaka-api/book"

	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"open":   "This is root directory",
		"status": "200 Success!",
	})
}

func RootHandler2(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"open":   "This is root directory versi 2",
		"status": "200 Success!",
	})
}

func BooksHandler(c *gin.Context) {

	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func PostInputBooks(c *gin.Context) {
	var inputBooks book.InputBooks

	err := c.ShouldBindJSON(&inputBooks)
	if err != nil {

		// errorMessages := []string{}
		// for_, e := range err.{validator.ValidationErrors} {
		// 		errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
		// 		errorMessages = append(errorerrorMessages, errerrorMessage	)
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"errors": errorMessages
		// 	})
		// 	return
		// }

		c.JSON(http.StatusOK, gin.H{
			"error": "Oops!, Something wrong",
		})
	} else {

		c.JSON(http.StatusOK, gin.H{
			"title": inputBooks.Title,
			"price": inputBooks.Price,
		})
	}
}
