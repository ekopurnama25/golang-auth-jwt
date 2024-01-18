package routes

import (
	"github.com/gofiber/fiber/v2"
)
func Setup(app *fiber.App) {
	SetupPublicAuthRoutes(app)
	//app.Use(middlewares.IsUserAuthenticated) 
	SetupAuthRoutes(app)
	SetupUserRoutes(app)
	SetupRoleRoutes(app)
	SetupCoffeRoutes(app)
}