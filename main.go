package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// Middleware: Logger
	app.Use(logger.New())

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Paths to your SSL certificate and private key files
	certFile := "./certs/server.crt"
	keyFile := "./certs/private.key"

	// Start the HTTPS server
	err := app.ListenTLS(":443", certFile, keyFile)
	if err != nil {
		panic("Error starting server: " + err.Error())
	}
}
