package routes

import (
	"golang-auth-apiweb-coffee/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthUsersRoutes(app *fiber.App) {
	app.Post("/api/auth", controllers.AUthUsersMiddlaware)
	app.Get("/api/home/", controllers.GetUsersLogin)
}