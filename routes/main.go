package routes

import "github.com/gofiber/fiber/v2"

func SetAllRoutes(app *fiber.App) {
	setUserRoutes(app)
	SetPostsRoutes(app)
	setOrgRoutes(app)
}
