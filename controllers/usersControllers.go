package controllers

import (
	"golang-auth-apiweb-coffee/database"
	"golang-auth-apiweb-coffee/models"
	"github.com/gofiber/fiber/v2"
)

func AllUsers(c *fiber.Ctx) error {
	var users []models.Users
	database.DB.Find(&users)
	return c.JSON(users)
}

func SaveUsers(c *fiber.Ctx) error {
	var user models.Users
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	result := database.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return c.JSON(user)
}