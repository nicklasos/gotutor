package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("db/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
}

func NewDB() *gorm.DB {
	return db
}
