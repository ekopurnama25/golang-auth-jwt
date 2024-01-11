package middlewares

import (
	"golang-auth-apiweb-coffee/database"
	"golang-auth-apiweb-coffee/models"
	"golang-auth-apiweb-coffee/util"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func IsUserAuthenticated(ctx *fiber.Ctx) error {
	tokenString := ctx.Get("Authorization")
	token := ctx.Cookies(util.CookieName)

	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	} else if token != "" {
		tokenString = ctx.Cookies(util.CookieName)
	}

	if tokenString == "" {
		return ctx.JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	// if tokenString != true {
	// 	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": 401, "message": "Yoor Token Expired"})
	// }

	Token, err := util.ParseToken(tokenString);
	if err != nil {
		return err
	}


	var user models.Users 
	database.DB.Where("id = ?",Token).First(&user)
	return ctx.Next()
}