package routes

import (
	"github.com/adrianetp/yconnect_backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetPostsRoutes(app *fiber.App) {
	app.Get("/posts/ig/getFromTag", controllers.GetPostsFromTag)
}
