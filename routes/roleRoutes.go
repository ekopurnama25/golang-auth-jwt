package routes

import (
	"golang-auth-apiweb-coffee/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoleRoutes(app *fiber.App) {
	app.Get("/api/roles", controllers.AllRoles)
	app.Get("/api/roles/:id", controllers.GetIdRole)
	app.Post("/api/roles", controllers.SaveRoles)
	app.Put("/api/roles/:id", controllers.UpdateRoleId)
	app.Delete("/api/roles/:id", controllers.DeleteIdRole)
}