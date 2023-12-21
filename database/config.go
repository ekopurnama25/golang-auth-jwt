package database

import (
	"golang-auth-apiweb-coffee/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:@tcp(127.0.0.1:3306)/golang-auth-apiweb-coffee?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}
	db.AutoMigrate(&models.AuthUserTokens{}, &models.Role{}, &models.Users{})
	DB = db
}