package controllers

import (
	"context"

	"github.com/adrianetp/yconnect_backend/config"
	"github.com/adrianetp/yconnect_backend/encryption"
	"github.com/adrianetp/yconnect_backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func Login(c *fiber.Ctx) error {
	var userBody models.UserLogin
	c.BodyParser(&userBody)

	res := config.Database.Collection("Users").FindOne(context.TODO(), bson.D{{"email", userBody.Email}})

	var user models.User

	res.Decode(&user)
	if user.Id.IsZero() {
		return c.JSON(fiber.Map{
			"status": 400,
			"error":  "no user found",
		})

	}
	pass, err := encryption.DecryptBase64(user.Password)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"error":  err.Error(),
		})
	}
	if pass != userBody.Password {
		return c.JSON(fiber.Map{
			"status": 400,
			"error":  "wrong password",
		})
	}

	return c.JSON(fiber.Map{
		"status": 200,
		"user":   user,
	})

}
