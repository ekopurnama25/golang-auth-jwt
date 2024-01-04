package middleware

import (
	"golang-auth-apiweb-coffee/database"
	"golang-auth-apiweb-coffee/models"
	"golang-auth-apiweb-coffee/util"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

const (
	DefaultHeaderAuththentication string = "Authorization"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func Authenticated(c *fiber.Ctx) error {
	headerKey := c.Get(DefaultHeaderAuththentication)
	token := c.Cookies(util.CookieName)
	if headerKey == "" {
		return c.JSON(fiber.Map{
			"message": "headerKey not Null",
		})
	}
	if token == "" { 
		return c.JSON(fiber.Map{
			"message": "token not Null",
		})
	}
	if headerKey != "" && headerKey == token {
		var ENV = goDotEnvVariable("TOKEN_SCRET")
		Token, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(ENV), nil
		})
		if err != nil {
			return err
		}
		Isclaims:= Token.Claims.(*jwt.StandardClaims)
		var users models.Users 
		result := database.DB.Where("user_id = ?",Isclaims).First(&users)
		if result.RowsAffected > 0 {
			return c.JSON(true)
		}else{
			return c.JSON(false)
		}
	}
	return c.JSON(fiber.Map{
			"message": "false",
		})
}