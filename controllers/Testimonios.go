package controllers

import (
	"context"
	"time"

	"github.com/adrianetp/yconnect_backend/config"
	"github.com/adrianetp/yconnect_backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddTestimonio(c *fiber.Ctx) error {
	var body struct {
		Testimonio models.Testimonio `json:testimonio`
		Token      string            `json:token`
	}
	body.Testimonio.Id = primitive.NewObjectID()
	body.Testimonio.TimeStamp = primitive.NewDateTimeFromTime(time.Now())
	c.BodyParser(&body)
	_, err := validateToken(body.Token)
	if err != nil {
		return c.JSON(fiber.Map{
			"status": 400,
			"error":  "invalid token",
		})
	}
	res, err := config.Database.Collection("Testimonios").InsertOne(context.TODO(), body.Testimonio)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"error":  err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status": 200,
		"data":   res,
	})
}

func GetTestimoniosByOrg(c *fiber.Ctx) error {
	var body struct {
		OrgId primitive.ObjectID `json:orgId`
		Token string             `json:token`
	}

	c.BodyParser(&body)
	_, err := validateToken(body.Token)
	if err != nil {
		return c.JSON(fiber.Map{
			"status": 400,
			"error":  "invalid token",
		})
	}

	resO := config.Database.Collection("Organizations").FindOne(context.TODO(), bson.D{{"_id", body.OrgId}})

	var org models.Organization

	resO.Decode(&org)

	if org.ID.IsZero() {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"error":  "organization not found",
		})
	}
	res, err := config.Database.Collection("Testimonios").Find(context.TODO(), bson.D{{"orgId", body.OrgId}})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"error":  err.Error(),
		})
	}
	var Testimonios []models.Testimonio
	for res.Next(context.TODO()) {
		var testimonio models.Testimonio
		res.Decode(&testimonio)
		Testimonios = append(Testimonios, testimonio)
	}
	return c.Status(200).JSON(fiber.Map{
		"status": 200,
		"forums": Testimonios,
	})
}

func GradeTestimonios(c *fiber.Ctx) error {
	var body struct {
		TestimonioId primitive.ObjectID `json:testimonioId`
		Grade        int                `json:grade`
		Token        string             `json:token`
	}

	c.BodyParser(&body)
	_, err := validateToken(body.Token)
	if err != nil {
		return c.JSON(fiber.Map{
			"status": 400,
			"error":  "invalid token",
		})
	}

	res, err := config.Database.Collection("Testimonios").UpdateOne(context.TODO(), bson.D{{"_id", body.TestimonioId}}, bson.D{{"$set", bson.D{{"grade", body.Grade}}}})

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"error":  err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": 200,
		"result": res,
	})

}
