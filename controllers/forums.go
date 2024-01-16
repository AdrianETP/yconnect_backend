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

func AddForum(c *fiber.Ctx) error {
	var body models.Forums
	body.Id = primitive.NewObjectID()
	body.TimeStamp = primitive.NewDateTimeFromTime(time.Now())
	c.BodyParser(&body)
	// verificamos que el id del usuario y de la organizacion sean existentes
	// usuario:
	resU := config.Database.Collection("Users").FindOne(context.TODO(), bson.D{{"_id", body.UserId}})

	var user models.User
	resU.Decode(&user)

	if user.Id.IsZero() {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"error":  "user doesn't exist",
		})
	}
	resO := config.Database.Collection("Organizations").FindOne(context.TODO(), bson.D{{"_id", body.OrgId}})

	var Org models.Organization

	resO.Decode(&Org)

	if Org.ID.IsZero() {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"error":  "Organization doesn't exist",
		})

	}
	res, err := config.Database.Collection("Forums").InsertOne(context.TODO(), body)

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
