package database

import (
	"golang-auth-apiweb-coffee/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "user:pass@protocol(127.0.0.1:3306)golang-auth-apiweb-coffee/?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}
	db.AutoMigrate(&models.Permission{}, &models.Role{}, &models.User{})
	DB = db
}