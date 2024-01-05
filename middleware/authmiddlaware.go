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
	authHeader := c.Get("Authorization")
	cookie := c.Cookies(util.CookieName)
	if authHeader == "" {
		return c.JSON(fiber.Map{
			"message": "Authorization token is required",
		})
	}
	if cookie == "" { 
		return c.JSON(fiber.Map{
			"message": "token not Null",
		})
	}
	

	if authHeader != "" && authHeader == cookie {
		var ENV = goDotEnvVariable("TOKEN_SCRET")
		Token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(ENV), nil
		})

		if err != nil {
			return err
		}
		
		Isclaims:= Token.Claims.(*jwt.StandardClaims)
		var user models.Users 
		result := database.DB.Where("id = ?",Isclaims.Issuer).First(&user)
		if result.RowsAffected > 0 {
				return c.Next()
		}else{
				return c.JSON(fiber.Map{
				"message": "false",
			})
		}
	}
	return c.JSON(fiber.Map{
		"message": "false",
	})
}