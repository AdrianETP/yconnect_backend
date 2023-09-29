package controllers

import (
	"context"

	"github.com/adrianetp/yconnect_backend/config"
	"github.com/adrianetp/yconnect_backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateOrganization(c *fiber.Ctx) error {
	var Org models.Organization
	c.BodyParser(&Org)
	result, error := config.Database.Collection("organization").InsertOne(context.TODO(), Org)

	if error != nil {
		c.JSON(fiber.Map{
			"status": 400,
			"error":  error,
		})
	}

	return c.JSON(fiber.Map{
		"status": 200,
		"data":   result,
	})
}

func SearchOrg(c *fiber.Ctx) error {
	var Tag models.TagType
	c.BodyParser(&Tag)
	tags := Tag.Tags

	orgCol := config.Database.Collection("organization")

	var organizations []models.Organization
	for _, t := range tags {

		results, _ := orgCol.Find(context.TODO(), bson.D{
			{"tags", bson.D{{"$elemMatch", bson.D{{"$eq", t}}}}},
		})
        var newOrgs []models.Organization

		for results.Next(context.TODO()) {
			var organization models.Organization
			results.Decode(&organization)
			organizations = append(newOrgs, organization)

		}
        organizations = append(organizations, newOrgs...)
	}
	return c.JSON(fiber.Map{
		"status":        200,
		"organizations": organizations,
	})

}

func GetAllOrgs(c *fiber.Ctx) error {
	// vamos a guardar los usuarios decodificados aqui
	var users []models.Organization
	// aqui vamos a llamar a mongo y decirle que encuentre a usuarios pero sin filtro ( osea que saque a todos los usuarios)
	results, err := config.Database.Collection("organization").Find(context.TODO(), bson.M{})
	if err != nil {
		return c.JSON(fiber.Map{
			"status": 400,
			"error":  err,
		})
	}
	// aca vamos a iterar por todos los resultados y decodificarlos
	for results.Next(context.TODO()) {
		var user models.Organization
		results.Decode(&user)
		users = append(users, user)
	}
	// regresamos a los usuarios como json
	return c.JSON(fiber.Map{
		"status": 200,
		"data":   users,
	})

}
