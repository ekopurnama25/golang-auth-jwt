package routes

import (
	"golang-auth-apiweb-coffee/middleware"

	"github.com/gofiber/fiber/v2"
)
func Setup(app *fiber.App) {
	SetupPublicAuthRoutes(app)
	app.Use(middleware.Authenticated) 
	SetupAuthRoutes(app)
	SetupUserRoutes(app)
	SetupRoleRoutes(app)

}