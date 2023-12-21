package controllers

import (
	"golang-auth-apiweb-coffee/database"
	"golang-auth-apiweb-coffee/models"
	"github.com/gofiber/fiber/v2"
)

func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role
	database.DB.Find(&roles)
	return c.JSON(roles)
}

func SaveRoles(c *fiber.Ctx) error {
	var role models.Role
	if err := c.BodyParser(&role); err != nil {
		return err
	}
	result := database.DB.Create(&role)
	if result.Error != nil {
		return result.Error
	}
	return c.JSON(role)
}