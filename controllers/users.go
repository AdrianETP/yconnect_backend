package controllers

import (
	"context"

	"github.com/adrianetp/yconnect_backend/config"
	"github.com/adrianetp/yconnect_backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUser(c *fiber.Ctx) error {
	var user models.User
	c.BodyParser(&user)
	result, err := config.Database.Collection("Users").InsertOne(context.TODO(), user)
	if err != nil {
		return c.JSON(fiber.Map{
			"status": 400,
			"error":  err,
		})
	}

	return c.JSON(fiber.Map{
		"status": 200,
		"data":   result,
	})
}

// vamos a obtener a todos los usuarios en esta llamada get
func GetAllUsers(c *fiber.Ctx) error {
	// vamos a guardar los usuarios decodificados aqui
	var users []models.User
	// aqui vamos a llamar a mongo y decirle que encuentre a usuarios pero sin filtro ( osea que saque a todos los usuarios)
	results, err := config.Database.Collection("Users").Find(context.TODO(), bson.M{})
	if err != nil {
		return c.JSON(fiber.Map{
			"status": 400,
			"error":  err,
		})
	}
	// aca vamos a iterar por todos los resultados y decodificarlos
	for results.Next(context.TODO()) {
		var user models.User
		results.Decode(&user)
		users = append(users, user)
	}
	// regresamos a los usuarios como json
	return c.JSON(fiber.Map{
		"status": 200,
		"data":   users,
	})
}

func AddtoFavorites(c *fiber.Ctx) error {
	var body struct {
		User         primitive.ObjectID `json:user`
		Organization primitive.ObjectID `json:organization`
	}

	c.BodyParser(&body)

	results, err := config.Database.Collection("Users").
		UpdateOne(context.TODO(), bson.D{{"_id", body.User}}, bson.D{{"$push", bson.D{{"favorites", body.Organization}}}})
	if err != nil {
		return c.JSON(fiber.Map{
			"status": 400,
			"error":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": 200,
		"result": results,
	})
}

func AddTags(c *fiber.Ctx) error {
	var body struct {
		UserId primitive.ObjectID `json:userid`
		Tags   []string           `json:tags`
	}

	for _, v := range body.Tags {
		_, err := config.Database.Collection("Users").
			UpdateOne(context.TODO(), bson.D{{"_id", body.UserId}}, bson.D{{"$push", bson.D{{"tags", v}}}})
		if err != nil {
			return c.JSON(fiber.Map{
				"status": 400,
				"error":  err.Error(),
			})
		}
	}
	return c.JSON(fiber.Map{
		"status": 200,
		"data":   body.Tags,
	})
}
