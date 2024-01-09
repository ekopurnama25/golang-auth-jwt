package middlewares

import (
	"golang-auth-apiweb-coffee/database"
	"golang-auth-apiweb-coffee/models"
	"golang-auth-apiweb-coffee/util"

	"github.com/gofiber/fiber/v2"
)

func IsUserAuthenticated(ctx *fiber.Ctx) error {
	tokenString := ctx.Get("Authorization")
	token := ctx.Cookies(util.CookieName)

	if tokenString == token {
		Token, err := util.ParseToken(tokenString);
		if err != nil {
			return err
		}
		var user models.Users 
		result := database.DB.Where("id = ?",Token).First(&user)
		if result.RowsAffected > 0 {
				return ctx.Next()
		}else{
				return ctx.JSON(fiber.Map{
				"message": "false",
			})
		}
	}
	return ctx.Next()
}