package routes

import (
	"github.com/adrianetp/yconnect_backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetAuthRoutes(app *fiber.App) {

	app.Post("/auth/login", controllers.Login)
}
