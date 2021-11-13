package controller

import (
	"gin_test/model"
	"gin_test/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BookAdd(c *gin.Context) {
	book := model.Book{}
	err := c.Bind(&book)
	if err == nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	bookService := service.BookService{}
	err = bookService.SetBook(&book)
	if err == nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func BookList(c *gin.Context) {
	c.JSONP(http.StatusOK, gin.H{
		"message": "this is list",
	})
}

func BookUpdate(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"status": "this is update",
	})
}

func BookDelete(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"status": "this is delete",
	})
}
