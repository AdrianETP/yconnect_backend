package routes

import (
	"github.com/adrianetp/yconnect_backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func setOrgRoutes(app *fiber.App) {
	app.Post("/organizations", controllers.CreateOrganization)
	app.Get("/organizations", controllers.GetAllOrgs)
	app.Post("/organizations/searchByTag", controllers.GetOrgByTag)
	app.Post("/organizations/Favorites", controllers.GetFavorites)
	app.Post("/organizations/SearchByName", controllers.GetOrgByName)
	app.Post("/organizations/Delete", controllers.DeleteOrg)
	app.Post("/organizations/SearchById", controllers.GetOrgById)
	app.Post("/organizations/ModifyOrg", controllers.ModifyOrg)
	app.Post("/organizations/SendMail", controllers.SendMail)
}
