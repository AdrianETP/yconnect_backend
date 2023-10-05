package main

import (
	"log"
	"os"

	"github.com/adrianetp/yconnect_backend/config"
	"github.com/adrianetp/yconnect_backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	config.ConnectDB()
	port := os.Getenv("PORT")
	app := fiber.New()

	routes.SetAllRoutes(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	app.Listen(":" + port)
}
