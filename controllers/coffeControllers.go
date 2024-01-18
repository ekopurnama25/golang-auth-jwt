package controllers

import (
	"fmt"
	"golang-auth-apiweb-coffee/database"
	"golang-auth-apiweb-coffee/models"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CoffeGetAll(c *fiber.Ctx) error {
	var coffe []models.Coffe
	database.DB.Find(&coffe)
	return c.JSON(fiber.Map{
		"message": "Berhasil Menampilkan Data Coffe",
		"coffe": coffe,
	})
}

func CoffePost(c * fiber.Ctx) error {
	file, err := c.FormFile("image")
    if err != nil {
        log.Println("image upload error --> ", err)
        return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
    }
	CoffeTExt := struct {
		JenisCoffe string `json:"jenis_coffe"`
		HargaCoffe    string `json:"harga_coffe"`
		DescriptionCoffe string `json:"description_coffe"`
	}{}
	if err := c.BodyParser(&CoffeTExt); err != nil {
		return err
	}
    uniqueId := uuid.New()
    filename := strings.Replace(uniqueId.String(), "-", "", -1)
    fileExt := strings.Split(file.Filename, ".")[1]
    image := fmt.Sprintf("%s.%s", filename, fileExt)
    err = c.SaveFile(file, fmt.Sprintf("./util/img_coffe/%s", image))
    if err != nil {
        log.Println("image save error --> ", err)
        return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
    }
    imageUrl := fmt.Sprintf("http://localhost:5000/util/img_coffe/%s", image)
    data := models.Coffe { 
		JenisCoffe : CoffeTExt.JenisCoffe,
		HargaCoffe :  CoffeTExt.HargaCoffe,
		ImagesCoffe: image,
		UrlImageCoffe: imageUrl,
		DescriptionCoffe : CoffeTExt.DescriptionCoffe,
    }
	result := database.DB.Create(&data)
	if result.Error != nil {
		return c.JSON(fiber.Map{
		"message": "Gagal Menambah Data",
	})
	}
	return c.JSON(fiber.Map{
		"message": "Berhasil Menambah Data",
		"data": data,
	})
}