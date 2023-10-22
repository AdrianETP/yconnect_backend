package main

import (
	"log"
	"os"

	"github.com/adrianetp/yconnect_backend/config"
	"github.com/adrianetp/yconnect_backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	// Middleware: Logger
	app.Use(logger.New())

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Paths to your SSL certificate and private key files
	certFile := os.Getenv("CERTPATH")
	keyFile := os.Getenv("KEYPATH")
	config.ConnectDB()
	routes.SetAllRoutes(app)

	// Start the HTTPS server
	err := app.ListenTLS(":443", certFile, keyFile)
	if err != nil {
		panic("Error starting server: " + err.Error())
	}
}
