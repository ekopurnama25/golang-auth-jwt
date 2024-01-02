package controllers

import (
	"golang-auth-apiweb-coffee/database"
	"golang-auth-apiweb-coffee/models"

	"github.com/gofiber/fiber/v2"
)

func AllUsers(c *fiber.Ctx) error {
	var users []models.Users
	database.DB.Preload("Role").Find(&users)
	return c.JSON(fiber.Map{
		"message": "Berhasil Menampilkan Data",
		"data": users,
	})
}

func SaveUsers(c *fiber.Ctx) error {
	var users models.Users
	if err := c.BodyParser(&users); err != nil {
		return err
	}
	const DefaultPasswordForNewUsers = "1234"
	users.SetPassword(DefaultPasswordForNewUsers)
	result := database.DB.Create(&users)
	if result.Error != nil {
		return c.JSON(fiber.Map{
		"message": "Gagal Menambah Data",
	})
	}
	return c.JSON(fiber.Map{
		"message": "Berhasil Menambah Data",
		"data": users,
	})
}

func GetIdUsers(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	var users models.Users
	err = database.DB.Preload("Role").First(&users, userId).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound)
	}
	return c.JSON(fiber.Map{
		"message": "Berhasil Menambah Data",
		"data": users,
	})
}

func DeleteIdUsers(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	result := database.DB.Delete(&models.Users{}, userId)
	if result.Error != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusNotFound)
	}
	return c.JSON(fiber.Map{
		"message": "Berhasil Menghapus Data",
		"id_users": userId,
	})
}

func UpdateUsers(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	var users models.Users
	if err := c.BodyParser(&users); err != nil {
		return err
	}
	users.Id = uint(userId)
	result := database.DB.Model(&users).Updates(users)
	if result.Error != nil {
		return err
	}
	return c.JSON(users)
}