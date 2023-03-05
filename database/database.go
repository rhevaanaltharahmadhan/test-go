package database

import (
	"coding-test/helpers"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {

	var err error

	env := godotenv.Load()

	if env != nil {
		helpers.Logger("error", "Error getting env")
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbPort, dbName)

	db, err = gorm.Open(mysql.Open(dbURI), &gorm.Config{})

	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
	}
	fmt.Println("Success | connected db MySQL")
	fmt.Println(db)

	// db.Debug().AutoMigrate(entities.Customer{}, entities.Limit{}, entities.Transaction{})

}

func GetDB() *gorm.DB {
	return db
}
