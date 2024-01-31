package routes

import (
	"github.com/adrianetp/yconnect_backend/controllers"
	"github.com/gofiber/fiber/v2"
)

// linkeamos un controller con una ruta (esta funcion se llama en el main del folder routes)
func setOrgRoutes(app *fiber.App) {
	app.Post("/organizations", controllers.CreateOrganization)
	app.Post("/organizations/getAll", controllers.GetAllOrgs)
	app.Post("/organizations/searchByTag", controllers.GetOrgByTag)
	app.Post("/organizations/Favorites", controllers.GetFavorites)
	app.Post("/organizations/SearchByName", controllers.GetOrgByName)
	app.Post("/organizations/Delete", controllers.DeleteOrg)
	app.Post("/organizations/SearchById", controllers.GetOrgById)
	app.Post("/organizations/ModifyOrg", controllers.ModifyOrg)
	app.Post("/organizations/SendMail", controllers.SendMail)
	app.Post("/organizations/ChangeGrade", controllers.ChangeGrade)
}
