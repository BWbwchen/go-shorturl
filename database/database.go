package database

import (
	"fmt"
	"os"
	. "shorturl/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB global database
var DB *gorm.DB

func init() {
	// load environment variable
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	HOST := os.Getenv("HOST")
	DBUSER := os.Getenv("DBUSER")
	PASSWORD := os.Getenv("PASSWORD")
	DBNAME := os.Getenv("DBNAME")
	PORT := os.Getenv("PORT")

	var dsn string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei", HOST, DBUSER, PASSWORD, DBNAME, PORT)

	bb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB = bb

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
		panic(err)
	}

	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
		panic(DB.Error)
	}
}

func Insert(t ShorturlSturct) {
	result := DB.Select("shortname", "url").Create(&t)
	if result.Error != nil {
		panic(result.Error)
	}
}

func Find(shortName string) (string, int) {
	var t ShorturlSturct
	result := DB.Where("shortName = ?", shortName).First(&t)
	if result.Error != nil {
		// not found
		return "link not found", NotFound
	}
	return t.URL, Success
}
