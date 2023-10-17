package routes

import (
	"github.com/adrianetp/yconnect_backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func setUserRoutes(app *fiber.App) {
	app.Get("/users", controllers.GetAllUsers)
	app.Post("/users", controllers.AddUser)
	app.Post("/users/addFavorites", controllers.AddtoFavorites)
	app.Post("/users/addTags", controllers.AddTags)
	app.Post("/users/Delete", controllers.DeleteUser)
	app.Post("/users/Update", controllers.EditUser)
}
