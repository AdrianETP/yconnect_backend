package routes

import (
	"github.com/adrianetp/yconnect_backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func setUserRoutes(app *fiber.App){
    app.Get("/users" , controllers.GetAllUsers)
    app.Post("/users", controllers.AddUser)
}
