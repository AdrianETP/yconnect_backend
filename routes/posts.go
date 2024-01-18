package routes

import (
	"github.com/adrianetp/yconnect_backend/controllers"
	"github.com/gofiber/fiber/v2"
)

// linkeamos un controller con una ruta (esta funcion se llama en el main del folder routes)
func setPostRoutes(app *fiber.App) {
	app.Get("/posts", controllers.GetPosts)
	app.Post("/posts", controllers.AddPost)
	app.Post("/posts/addlike", controllers.AddLike)
	app.Post("/posts/addComment", controllers.AddComment)
}
