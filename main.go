package main

import (
	"golang-auth-apiweb-coffee/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()

    app := fiber.New()

	// Add CORS Middleware so the frontend get the cookie
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	//routes.Setup(app)

	app.Listen(":8000")
}