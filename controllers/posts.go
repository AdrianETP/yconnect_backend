package controllers

import (
	"context"
	"time"
	"fmt"

	"github.com/adrianetp/yconnect_backend/config"
	"github.com/adrianetp/yconnect_backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// funcion para agregar un publicacion
func AddPost(c *fiber.Ctx) error {
	// variable para parsear el body de la request
	var post models.Post
	// parseamos el body de la request

	c.BodyParser(&post)

	// guardamos todos los valores predeterminados
	post.Id = primitive.NewObjectID()
	post.Likes = []primitive.ObjectID{}
	post.Comments = []models.Comment{}
	post.TimeStamp = primitive.NewDateTimeFromTime(time.Now())

	if post.MediaUrls == nil {
		post.MediaUrls = []string{}
	}

	// revisar que organizacion existe
	resO := config.Database.Collection("Organizations").FindOne(context.TODO(), bson.D{{"_id", post.OrgId}})

	var Org models.Organization

	resO.Decode(&Org)

	if Org.ID.IsZero() {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"error":  "Organization doesn't exist",
		})

	}

	result, err := config.Database.Collection("Posts").InsertOne(context.TODO(), post)
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

// lista de publicacion de una organizacion
func GetPosts(c *fiber.Ctx) error {
	var body struct {
		OrgId primitive.ObjectID `json:orgId`
	}

	c.BodyParser(&body)

	// revissar que la organizacion existe
	resO := config.Database.Collection("Organizations").FindOne(context.TODO(), bson.D{{"_id", body.OrgId}})

	var org models.Organization

	resO.Decode(&org)

	if org.ID.IsZero() {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"error":  "organization not found",
		})
	}
	res, err := config.Database.Collection("Posts").Find(context.TODO(), bson.D{{"orgId", body.OrgId}})

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"error":  err.Error(),
		})
	}
	var posts []models.Post
	for res.Next(context.TODO()) {
		var post models.Post
		res.Decode(&post)
		posts = append(posts, post)
	}
	return c.Status(200).JSON(fiber.Map{
		"status": 200,
		"posts": posts,
	})
}


// funcion para agregar likes a una publicacion
func AddLike(c *fiber.Ctx) error {
	// variable para parsear el body
	var body struct {
		UserID primitive.ObjectID `json:userId`
		PostID primitive.ObjectID `json:postId`
	}

	// parseamos el body
	c.BodyParser(&body)


	// verificamos que el usuario existe
	resU := config.Database.Collection("Users").FindOne(context.TODO(), bson.D{{"_id", body.UserID}})

	var user models.User
	resU.Decode(&user)

	if user.Id.IsZero() {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"error":  "user doesn't exist",
		})
	}

	results, err := config.Database.Collection("Posts").
		UpdateOne(context.TODO(), bson.D{{"_id", body.PostID}}, bson.D{{"$addToSet", bson.D{{"likes", body.UserID}}}})

	// si el valor es nulo, haz que sea una lista con el id del usuario

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
		"data":   results,
	})
}

// funcion para agregar comentarios a una publicacion
func AddComment(c *fiber.Ctx) error {
	// variable para parsear el comment
	var comment models.Comment

	var postId struct {
		PostId primitive.ObjectID `json:postId`
	}

	// parseamos el body
	c.BodyParser(&postId)
	c.BodyParser(&comment)

	comment.TimeStamp = primitive.NewDateTimeFromTime(time.Now())

	fmt.Println(postId)
	fmt.Println(comment)

	resU := config.Database.Collection("Users").FindOne(context.TODO(), bson.D{{"_id", comment.UserID}})

	var user models.User
	resU.Decode(&user)

	if user.Id.IsZero() {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"error":  "user doesn't exist",
		})
	}

	// agregamos el id de la organizacion a los favoritos del usuario
	results, err := config.Database.Collection("Posts").
		UpdateOne(context.TODO(), bson.D{{"_id", postId.PostId}}, bson.D{{"$push", bson.D{{"comments", comment}}}})

	// si el valor es nulo, haz que sea una lista con el id del usuario

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
		"data":   results,
	})
}

// get user liked
