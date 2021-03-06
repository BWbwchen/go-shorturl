package database

import (
	"fmt"
	. "shorturl/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB global database
var DB *gorm.DB
var dsn string = "host=host user=user password=password dbname=dbname port=port sslmode=disable TimeZone=Asia/Taipei"

func init() {
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
