package routes

import (
	"golang-auth-apiweb-coffee/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupCoffeRoutes(app *fiber.App) {
	app.Get("/api/users/", controllers.AllUsers)
}