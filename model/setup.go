package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("mysql", "book.db")

	if err != nil {
		panic("Failed to connect database!")
	}

	database.AutoMigrate(&Book{})

	DB = database
}
