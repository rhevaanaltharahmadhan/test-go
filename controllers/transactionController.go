package controllers

import (
	"coding-test/handlers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func C_CreateTransaction(c *gin.Context) {
	id := c.Param("id")
	customerId, _ := strconv.Atoi(id)

	data := &handlers.Transaction{}

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	response, err := data.H_CreateTransaction(customerId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Create Transaction Success",
		"data":    response,
	})
}
