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
	err := resU.Decode(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"error":  "Error parsing user",
		})
	}

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

func GetForumsFromOrg(c *fiber.Ctx) error {
	var body struct {
		OrgId primitive.ObjectID `json:orgId`
	}

	c.BodyParser(&body)

	resO := config.Database.Collection("Organizations").FindOne(context.TODO(), bson.D{{"_id", body.OrgId}})

	var org models.Organization

	resO.Decode(&org)

	if org.ID.IsZero() {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"error":  "organization not found",
		})
	}
	res, err := config.Database.Collection("Forums").Find(context.TODO(), bson.D{{"orgId", body.OrgId}})

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"error":  err.Error(),
		})
	}
	var forums []models.Forums
	for res.Next(context.TODO()) {
		var forum models.Forums
		res.Decode(&forum)
		forums = append(forums, forum)
	}
	return c.Status(200).JSON(fiber.Map{
		"status": 200,
		"forums": forums,
	})
}

func GetForumsByUser(c *fiber.Ctx) error {

	var body struct {
		UserId primitive.ObjectID `json:userId`
	}
	c.BodyParser(&body)

	// check if user id exists

	resU := config.Database.Collection("Users").FindOne(context.TODO(), bson.D{{"_id", body.UserId}})

	var user models.User

	resU.Decode(&user)

	if user.Id.IsZero() {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"error":  "user not found",
		})
	}

	res, err := config.Database.Collection("Forums").Find(context.TODO(), bson.D{{"userId", body.UserId}})

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"error":  err.Error(),
		})

	}

	var forums []models.Forums

	for res.Next(context.TODO()) {
		var forum models.Forums
		res.Decode(&forum)
		forums = append(forums, forum)
	}
	return c.Status(200).JSON(fiber.Map{
		"status": 200,
		"forums": forums,
	})
}
