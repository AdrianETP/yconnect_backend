package routes

import (
	"github.com/adrianetp/yconnect_backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetTestimonioRoutes(app *fiber.App) {
	app.Post("/testimonios", controllers.AddTestimonio)
	app.Post("/testimonios/getByOrg", controllers.GetTestimoniosByOrg)
}
