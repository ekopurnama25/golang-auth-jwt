package routes

import (
	"golang-auth-apiweb-coffee/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupCoffeRoutes(app *fiber.App) {
	app.Get("/api/coffe/", controllers.CoffeGetAll)
	app.Post("/api/coffe/", controllers.CoffePost)
	app.Delete("/api/coffe/:id", controllers.DeleteCoffe)
	app.Get("/api/coffe/:id", controllers.GetIdCoffe)
	app.Put("/api/coffe/:id", controllers.UpdateDataCoffe)
}