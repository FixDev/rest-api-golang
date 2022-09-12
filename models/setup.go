package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_book?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect database!")
	}

	database.AutoMigrate(&Book{})

	DB = database
}
