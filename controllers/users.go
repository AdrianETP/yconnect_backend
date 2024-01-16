package controllers

import (
	"context"

	"github.com/adrianetp/yconnect_backend/config"
	"github.com/adrianetp/yconnect_backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// funcion para agregar un usuario
func AddUser(c *fiber.Ctx) error {
	// variable para parsear el body de la request
	var user models.User
	// parseamos el body de la request
	c.BodyParser(&user)
	// agregamos un id unico al usuario
	user.Id = primitive.NewObjectID()
	// agregamos el usuario a la base de datos
	result, err := config.Database.Collection("Users").InsertOne(context.TODO(), user)
	// si da un error
	if err != nil {
		// regresamos el error
		return c.JSON(fiber.Map{
			"status": 400,
			"error":  err,
		})
	}

	// regresamos el resultado
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

// agregamos organizaciones favoritas al usuario
func AddtoFavorites(c *fiber.Ctx) error {
	// variable para parsear el body
	var body struct {
		UserID         primitive.ObjectID `json:user`
		OrganizationID primitive.ObjectID `json:organization`
	}

	// parseamos el body
	c.BodyParser(&body)

	// agregamos el id de la organizacion a los favoritos del usuario
	results, err := config.Database.Collection("Users").
		UpdateOne(context.TODO(), bson.D{{"_id", body.UserID}}, bson.D{{"$push", bson.D{{"favorites", body.OrganizationID}}}})
		// si da un error
	if err != nil {
		// regresamos el error
		return c.JSON(fiber.Map{
			"status": 400,
			"error":  err.Error(),
		})
	}
	// regresamos el resultado
	return c.JSON(fiber.Map{
		"status": 200,
		"result": results,
	})
}

// funcion para modificar a un usuario
func ModifyUser(c *fiber.Ctx) error {
	// variable para parsear el body
	var body struct {
		User models.User `json:user`
	}
	// parseamos el body
	c.BodyParser(&body)
	// modificamos el
	result, err := config.Database.Collection("Users").
		UpdateOne(context.TODO(), bson.D{{"_id", body.User.Id}}, bson.D{
			{"name", body.User.Name},
			{"telephone", body.User.Telephone},
			{"tags", body.User.Tags},
			{"description", body.User.Description},
			{"favorites", body.User.Favorites},
			{"password", body.User.Password},
		})
		// si da un error
	if err != nil {
		// regresamos el error
		return c.JSON(fiber.Map{
			"error":  err.Error(),
			"status": 400,
		})
	}

	// regresamos los resultados
	return c.JSON(fiber.Map{
		"status": 200,
		"data":   result,
	})
}

// funcion para agregar tags de interes a un usuario
func AddTags(c *fiber.Ctx) error {
	// variable para parsear el body
	var body struct {
		UserId primitive.ObjectID `json:userid`
		Tags   []string           `json:tags`
	}

	// parseamos el body
	c.BodyParser(&body)

	// iteramos por cada tag
	for _, v := range body.Tags {
		// agregamos el tag a la lista de tags del usuario
		_, err := config.Database.Collection("Users").
			UpdateOne(context.TODO(), bson.D{{"_id", body.UserId}}, bson.D{{"$push", bson.D{{"tags", v}}}})
			// si hubo un error
		if err != nil {
			// regresamos el error
			return c.JSON(fiber.Map{
				"status": 400,
				"error":  err.Error(),
			})
		}
	}
	// regresamos el resultado
	return c.JSON(fiber.Map{
		"status": 200,
		"data":   body.Tags,
	})
}

// funcion para borrar un usuario
func DeleteUser(c *fiber.Ctx) error {
	// variable para parsear el body
	var body struct {
		UserId primitive.ObjectID `json:userid`
	}

	// parseamos el body
	c.BodyParser(&body)

	// borramos el usuario
	result, err := config.Database.Collection("Users").
		DeleteOne(context.TODO(), bson.D{{"_id", body.UserId}})
		// si hubo un error
	if err != nil {
		// regresamos el error
		return c.JSON(fiber.Map{
			"status": 400,
			"error":  err.Error(),
		})
	}

	// regresamos el resultado
	return c.JSON(fiber.Map{
		"status": 200,
		"data":   result,
	})
}

// funcion de login
// TODO crear una nueva funcion de login por que esta era dummie
func Login(c *fiber.Ctx) error {
	var body struct {
		Telephone string `json:telephone`
	}
	c.BodyParser(&body)

	result := config.Database.Collection("Users").FindOne(context.TODO(), bson.D{{"telephone", body.Telephone}})

	var user models.User
	result.Decode(&user)

	return c.JSON(fiber.Map{
		"status": 200,
		"result": user,
	})
}
