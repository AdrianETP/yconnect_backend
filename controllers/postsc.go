package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/adrianetp/yconnect_backend/config"
	"github.com/adrianetp/yconnect_backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetPostsFromTag(c *fiber.Ctx) error {
	c.Accepts("json", "text") // "json"
	c.Accepts("application/json")
	var body struct {
		Tags []string `json:tags`
	}
	c.BodyParser(&body)
	tags := body.Tags
	orgCol := config.Database.Collection("organization")

	results, _ := orgCol.Find(context.TODO(), bson.D{
		{"tags", bson.D{{"$elemMatch", bson.D{{"$eq", tags[0]}}}}},
	})

	i := 0
	var organizations []models.Organization

	for results.Next(context.TODO()) && i < 10 {
		var organization models.Organization
		results.Decode(&organization)
		organizations = append(organizations, organization)
		i++

	}
	// var posts []models.IgPost
	postRequest, err := http.NewRequest("GET", organizations[0].IgUrl, nil)
	if err != nil {
		return c.JSON(fiber.Map{
			"status": 400,
			"error":  err.Error(),
		})
	}

	client := &http.Client{}
	postResponse, err := client.Do(postRequest)
	if err != nil {
		return c.JSON(fiber.Map{
			"status": 400,
			"error":  err.Error(),
		})
	}

	defer postResponse.Body.Close()
	decoder := json.NewDecoder(postResponse.Body)
	post := &models.IgPost{}
	error := decoder.Decode(post)
	if error != nil {
		return c.JSON(fiber.Map{
			"status": 400,
			"error":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": 200,
		"data":   post,
		"url":    organizations[0].IgUrl,
	})
}

func GetPostsByName(c *fiber.Ctx) error {
	c.Accepts("json", "text") // "json"
	c.Accepts("application/json")
	var body struct {
		Name string
	}
	c.BodyParser(&body)
	orgCol := config.Database.Collection("organization")
	results, _ := orgCol.Find(context.TODO(), bson.D{
		{"name", body.Name},
	})

	i := 0
	var organizations []models.Organization

	for results.Next(context.TODO()) && i < 10 {
		var organization models.Organization
		results.Decode(&organization)
		organizations = append(organizations, organization)
		i++

	}
	for _, v := range organizations {
		print(v.Description)
	}

	return c.JSON(fiber.Map{
		"status": 200,
		"data":   "posts",
	})
}
