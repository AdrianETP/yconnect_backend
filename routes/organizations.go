package routes



import (
	"github.com/adrianetp/yconnect_backend/controllers"
	"github.com/gofiber/fiber/v2"
)


func setOrgRoutes(app *fiber.App){
    app.Post("/organizations", controllers.CreateOrganization)
}
