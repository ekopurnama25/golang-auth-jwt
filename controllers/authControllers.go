package controllers

import (
	"golang-auth-apiweb-coffee/database"
	"golang-auth-apiweb-coffee/models"
	"golang-auth-apiweb-coffee/util"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func AUthUsersMiddlaware(c *fiber.Ctx) error {
	loginDto := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&loginDto); err != nil {
		return err
	}

	var users models.Users
	result := database.DB.Where("email = ?",loginDto.Email).First(&users)
	if result.RowsAffected > 0 {
		 if err := bcrypt.CompareHashAndPassword(users.Password, []byte(loginDto.Password)); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{"message":"Incorrect password !"})
		}else{
			expirationTimeRefresh := time.Now().Add(24 * time.Hour)
			expirationTime := time.Now().Add(15 * time.Hour)
			Token, err := util.CreateToken(strconv.Itoa(int(users.Id)), expirationTime)
			Refresh, err := util.CreateRefreshToken(strconv.Itoa(int(users.Id)), expirationTimeRefresh)
			if err != nil {
				return err
			}

			var tokens []models.AuthUserTokens
			resultToken := database.DB.Where("user_id = ?", users.Id).First(&tokens)

			if resultToken.RowsAffected > 0 {
				tokenscreate := models.AuthUserTokens{
					AccessToken: Token,
					RefeshToken: Refresh,
					UserId: users.Id,
				}
				database.DB.Updates(&tokenscreate)
			}else{
				tokenscreate := models.AuthUserTokens{
					AccessToken: Token,
					RefeshToken: Refresh,
					UserId: users.Id,
				}
				database.DB.Create(&tokenscreate)
			}

			c.Cookie(&fiber.Cookie{
				Name:     util.CookieName,
				Value:    Token,
				Path:     "",
				Domain:   "",
				MaxAge:   0,
				Expires:  time.Now().Add(time.Hour * 15),
				Secure:   false,
				HTTPOnly: true,
				SameSite: "lax",
			})
			return c.JSON(fiber.Map{
				"AccessToken":Token,
				"RefreshToken":Refresh,
			})
		}
	}else{
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{"message":"Usename not found !"})
	}
}

func GetUsersLogin(c *fiber.Ctx) error {
	cookie := c.Cookies(util.CookieName)
	userId, err := util.ParseToken(cookie)
	if err != nil {
		return err
	}
	var users models.Users
	err = database.DB.Preload("Role").First(&users, userId).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound)
	}
	return c.JSON(users)
}