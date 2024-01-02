package controllers

import (
	"golang-auth-apiweb-coffee/database"
	"golang-auth-apiweb-coffee/models"

	"github.com/gofiber/fiber/v2"
)

func AUthUsersMiddlaware(c *fiber.Ctx) error {
	loginDto := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	var users []models.Users
	//print(loginDto);
	if err := c.BodyParser(&loginDto); err != nil {
		return err
	}
	result := database.DB.First(&users, loginDto.Email).Error

	return c.JSON(fiber.Map{
		"message": "successfully logged in",
		"data": result,
	})
}