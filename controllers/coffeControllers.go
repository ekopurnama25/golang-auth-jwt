package controllers

import (
	"golang-auth-apiweb-coffee/database"
	"golang-auth-apiweb-coffee/models"

	"github.com/gofiber/fiber/v2"
)

func CoffeGetAll(c *fiber.Ctx) error {
	var coffe []models.Coffe
	database.DB.Find(&coffe)
	return c.JSON(fiber.Map{
		"message": "Berhasil Menampilkan Data Coffe",
		"coffe": coffe,
	})
}