package main

import (
	"coding-test/controllers"
	"coding-test/database"
	"coding-test/helpers"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		helpers.Logger("error", "Error Getting Env")
	}

	database.GetDB()

	middleUrl := os.Getenv("MIDDLE_URL")

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "OPTIONS", "GET", "DELETE"},
		AllowHeaders: []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	}))

	// customer
	router.POST(middleUrl+"/customer/create", controllers.C_CreateCustomer)
	router.GET(middleUrl+"/customer/fetch", controllers.C_GetAllCustomer)

	// transaction
	router.POST(middleUrl+"/transaction/create/:id", controllers.C_CreateTransaction)

	port := os.Getenv("PORT")

	router.Run(":" + port)
}
