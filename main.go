package main

import (
	"golang-auth-apiweb-coffee/database"
	"golang-auth-apiweb-coffee/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {

	database.Connect()

    app := fiber.New()

	// Add CORS Middleware so the frontend get the cookie
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	

	routes.Setup(app)
	PORT:=goDotEnvVariable("APP_PORT")
	app.Listen(PORT)
}