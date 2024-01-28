package controllers

import (
	"errors"
	"fmt"
	"golang-auth-apiweb-coffee/database"
	"golang-auth-apiweb-coffee/models"
	"log"
	"os"
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

func DeleteCoffe(c * fiber.Ctx) error {
	CoffeId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	var coffe models.Coffe
	cekid := database.DB.First(&coffe, CoffeId).Error
	if cekid != nil {
		return fiber.NewError(fiber.StatusNotFound)
	}
	directory := "./util/img_coffe/"
	filename := directory + coffe.ImagesCoffe
	os.Remove(filename)
	result := database.DB.Delete(&coffe, CoffeId)
	if result.Error != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "Berhasil menghapus data",
		"status": 200,
		"id": CoffeId,
	});
}

func GetIdCoffe(c * fiber.Ctx) error {
	CoffeId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	var coffe models.Coffe
	err = database.DB.First(&coffe, CoffeId).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound)
	}
	return c.JSON(fiber.Map{
		"message": "Berhasil Menambah Data",
		"data": coffe,
	})
}

func UpdateDataCoffe(c *fiber.Ctx) error {
	CoffeId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
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
	if err := c.BodyParser(&CoffeTExt); err != nil {
		return err
	}
    uniqueId := uuid.New()
    filename := strings.Replace(uniqueId.String(), "-", "", -1)
    fileExt := strings.Split(file.Filename, ".")[1]
    image := fmt.Sprintf("%s.%s", filename, fileExt)
	if image  {
		save = c.SaveFile(file, fmt.Sprintf("./util/img_coffe/%s", image))
			if save != nil {
				log.Println("image save error --> ", err)
				return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
			}else{
				var coffe models.Coffe
				err = database.DB.First(&coffe, CoffeId).Error
				imageLama := coffe.ImagesCoffe;
				url := "./util/img_coffe/"
				e := os.Remove(url + imageLama)
				if e != nil { 
					log.Fatal(e) 
				}  
			}

			imageUrl := fmt.Sprintf("http://localhost:5000/util/img_coffe/%s", image)

			data := models.Coffe { 
				JenisCoffe : CoffeTExt.JenisCoffe,
				HargaCoffe :  CoffeTExt.HargaCoffe,
				ImagesCoffe: image,
				UrlImageCoffe: imageUrl,
				DescriptionCoffe : CoffeTExt.DescriptionCoffe,
			}

			var coffe models.Coffe
			coffe.Id = uint(CoffeId)
			result := database.DB.Model(&coffe).Updates(data)
			if result.Error != nil {
				return err
			}
			
		return c.JSON(fiber.Map{
			"Update": CoffeId,
			"status": data,
		})
	}else{
		data := models.Coffe { 
			JenisCoffe : CoffeTExt.JenisCoffe,
			HargaCoffe :  CoffeTExt.HargaCoffe,
			DescriptionCoffe : CoffeTExt.DescriptionCoffe,
    	}

		var coffe models.Coffe
		coffe.Id = uint(CoffeId)
		result := database.DB.Model(&coffe).Updates(data)
		if result.Error != nil {
			return err
		}
		
		return c.JSON(fiber.Map{
			"Update": CoffeId,
			"status": data,
		})
	}
}