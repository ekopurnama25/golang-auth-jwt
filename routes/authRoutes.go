package routes

import (
	"golang-auth-apiweb-coffee/controllers"
	"golang-auth-apiweb-coffee/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthUsersRoutes(app *fiber.App) {
	app.Post("/api/auth", controllers.AUthUsersMiddlaware)
	app.Get("/api/home/", middleware.Authenticated, controllers.GetUsersLogin)
}