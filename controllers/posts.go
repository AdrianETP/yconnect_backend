package controllers

import (
	"context"

	"github.com/adrianetp/yconnect_backend/config"
	"github.com/adrianetp/yconnect_backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetPostsFromTag(c *fiber.Ctx) error {
	c.Accepts("json", "text") // "json"
	c.Accepts("application/json")
	var posts []models.PostType
	var user models.User
	c.BodyParser(&user)
	tags := user.Tags

	orgCol := config.Database.Collection("organization")

	for _, t := range tags {

		results, _ := orgCol.Find(context.TODO(), bson.D{
			{"tags", bson.D{{"$elemMatch", bson.D{{"$eq", t}}}}},
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

			instaUser, err := config.Insta.VisitProfile(v.Igtag)
			if err != nil {
				return c.JSON(fiber.Map{
					"status": 400,
					"error":  err,
				})
			}
			for i, p := range instaUser.Feed.Items {
				if i >= 3 {
					break
				}
				var post models.PostType = models.PostType{
					User:    v.Igtag,
					Caption: p.Caption.Text,
					Image:   p.Images.GetBest(),
				}
				posts = append(posts, post)

			}

		}
	}

	return c.JSON(fiber.Map{
        "status":200,
		"data": posts,
	})
}
