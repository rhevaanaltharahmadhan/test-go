package controllers

import (
	"coding-test/handlers"
	"coding-test/helpers"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func C_CreateCustomer(c *gin.Context) {
	data := &handlers.Customer{}

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	response, err := data.H_CreateCustomer()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	Logger, _ := json.Marshal(&response)
	helpers.Logger("info", "Create Customer Success, Response: "+string(Logger))

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Create Customer Success",
		"data":    response,
	})
}

func C_GetAllCustomer(c *gin.Context) {

	response, err := handlers.H_GetAllCustomer()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Fetch Customers Success",
		"data":    response,
	})
}
