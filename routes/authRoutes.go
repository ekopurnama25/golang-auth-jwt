package routes

import (
	"golang-auth-apiweb-coffee/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupPublicAuthRoutes(app *fiber.App) {
	app.Post("/api/auth", controllers.AUthUsersMiddlaware)
}

func SetupAuthRoutes(app *fiber.App) {
	app.Get("/api/home/", controllers.GetUsersLogin)
}