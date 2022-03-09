package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	// "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDataBase() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Dbdriver := os.Getenv("DB_DRIVER")
	// DbHost := os.Getenv("DB_HOST")
	// DbUser := os.Getenv("DB_USER")
	// DbPassword := os.Getenv("DB_PASSWORD")
	// DbName := os.Getenv("DB_NAME")
	// DbPort := os.Getenv("DB_PORT")

	// DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	// DB, err = gorm.Open(Dbdriver, DBURL)

	// if err != nil {
	// 	fmt.Println("Cannot connect to database ", Dbdriver)
	// 	log.Fatal("connection error:", err)
	// } else {
	// 	fmt.Println("We are connected to the database ", Dbdriver)
	// }

	// DB.AutoMigrate(&User{})

	DB, err = gorm.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Println("Cannon connect to database")
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database")
	}

	DB.AutoMigrate(&User{})
}