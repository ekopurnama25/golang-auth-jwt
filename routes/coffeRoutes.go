package routes

import (
	"golang-auth-apiweb-coffee/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupCoffeRoutes(app *fiber.App) {
	app.Static("/utils/img_coffe", "./utils/img_coffe")
	app.Get("/api/coffe/", controllers.CoffeGetAll)
	app.Post("/api/coffe/", controllers.CoffePost)
}