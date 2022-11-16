package database

import (
	"example/login/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "root:@tcp(127.0.0.1:3306)/go_auth?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
	d.AutoMigrate(&models.User{})
}

func GetDB() *gorm.DB {
	return db
}
