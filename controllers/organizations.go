package controllers

import (
	"context"

	"github.com/adrianetp/yconnect_backend/config"
	"github.com/adrianetp/yconnect_backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GetOrgByTag(c *fiber.Ctx) error {
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
		"Organizations": organizations,
	})
}

func GetAllOrgs(c *fiber.Ctx) error {
	// vamos a guardar los usuarios decodificados aqui
	var organizations []models.Organization
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
		var organnization models.Organization
		results.Decode(&organnization)
		organizations = append(organizations, organnization)
	}
	// regresamos a los usuarios como json
	return c.JSON(fiber.Map{
		"status": 200,
		"data":   organizations,
	})
}

func GetFavorites(c *fiber.Ctx) error {
	var body struct {
		UserID primitive.ObjectID `json:userId`
	}
	c.BodyParser(&body)

	results := config.Database.Collection("Users").
		FindOne(context.TODO(), bson.D{{"_id", body.UserID}})
	var user models.User
	results.Decode(&user)
	var organizations []models.Organization
	for _, v := range user.Favorites {
		r, err := config.Database.Collection("organization").
			Find(context.TODO(), bson.D{{"_id", v}})
		if err != nil {
			return c.JSON(fiber.Map{
				"status": 400,
				"error":  err.Error(),
			})
		}

		for r.Next(context.TODO()) {
			var org models.Organization

			r.Decode(&org)
			organizations = append(organizations, org)

		}
	}

	return c.JSON(fiber.Map{
		"status": 200,
		"data":   organizations,
	})
}

func GetOrgByName(c *fiber.Ctx) error {
	var body struct {
		Name string `json:name`
	}
	c.BodyParser(&body)
	result, err := config.Database.Collection("organization").
		Find(context.TODO(), bson.D{{"name", bson.D{{"$regex", primitive.Regex{Pattern: ".*" + body.Name + ".*", Options: "i"}}}}})
	if err != nil {
		return c.JSON(fiber.Map{
			"status": 400,
			"error":  err.Error(),
		})
	}
	var organizations []models.Organization

	for result.Next(context.TODO()) {
		var organization models.Organization
		result.Decode(&organization)
		organizations = append(organizations, organization)
	}
	return c.JSON(fiber.Map{
		"status": 200,
		"data":   organizations,
	})
}

func GetOrgById(c *fiber.Ctx) error {
	var body struct {
		OrgId primitive.ObjectID `json:orgid`
	}

	c.BodyParser(&body)

	result := config.Database.Collection("organization").
		FindOne(context.TODO(), bson.D{{"_id", body.OrgId}})

	var organization models.Organization

	result.Decode(&organization)

	return c.JSON(fiber.Map{
		"status": 200,
		"data":   organization,
	})
}

func ModifyOrg(c *fiber.Ctx) error {
	var body struct {
		Organization models.Organization `json:organization`
	}

	result, err := config.Database.Collection("organization").
		UpdateOne(context.TODO(), bson.D{{"_id", body.Organization.ID}}, bson.D{
			{"name", body.Organization.Name},
			{"location", body.Organization.Location},
			{"telephone", body.Organization.Telephone},
			{"tags", body.Organization.Tags},
			{"igtag", body.Organization.Tags},
			{"igurrl", body.Organization.IgUrl},
			{"description", body.Organization.Description},
			{"email", body.Organization.Email},
		})
	if err != nil {
		return c.JSON(fiber.Map{
			"error":  err.Error(),
			"status": 400,
		})
	}

	return c.JSON(fiber.Map{
		"status": 200,
		"data":   result,
	})
}
