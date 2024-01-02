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

func GetIdRole(c *fiber.Ctx) error {
	roleId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	var role models.Role
	err = database.DB.First(&role, roleId).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound)
	}
	return c.JSON(fiber.Map{
		"message": "Berhasil Menambah Data",
		"data": role,
	})
}

func DeleteIdRole(c *fiber.Ctx) error {
	roleId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	result := database.DB.Delete(&models.Role{}, roleId)
	if result.Error != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusNotFound)
	}
	return c.JSON(fiber.Map{
		"message": "Berhasil Menghapus Data",
		"id_roles": roleId,
	})
}

func UpdateRoleId(c *fiber.Ctx) error {
	roleId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	var roles models.Role
	if err := c.BodyParser(&roles); err != nil {
		return err
	}
	roles.Id = uint(roleId)
	result := database.DB.Model(&roles).Updates(roles)
	if result.Error != nil {
		return err
	}
	return c.JSON(roles)
}