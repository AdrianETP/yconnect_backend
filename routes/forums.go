package routes

import (
	"github.com/adrianetp/yconnect_backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetForumRoutes(app *fiber.App) {
	app.Post("/forums", controllers.AddForum)
	app.Post("/forums/GetFromOrg", controllers.GetForumsFromOrg)
	app.Post("/forums/GetFromUser", controllers.GetForumsByUser)
}
