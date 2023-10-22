package controllers

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	"github.com/adrianetp/yconnect_backend/config"
	"github.com/adrianetp/yconnect_backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/gomail.v2"
)

func CreateOrganization(c *fiber.Ctx) error {
	var Org models.Organization
	Org.ID = primitive.NewObjectID()
	err := c.BodyParser(&Org)
	if err != nil {
		return c.JSON(fiber.Map{
			"status": 400,
			"error":  err.Error(),
		})
	}
	print(Org.Name)
	result, error := config.Database.Collection("organization").InsertOne(context.TODO(), Org)

	if error != nil {
		return c.JSON(fiber.Map{
			"status": 400,
			"error":  error.Error(),
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

		for results.Next(context.TODO()) {
			var organization models.Organization
			results.Decode(&organization)
			inOrg := false
			if len(organizations) <= 0 {
				organizations = append(organizations, organization)
			} else {
				for i := 0; i < len(organizations); i++ {
					if organization.ID.String() == organizations[i].ID.String() {
						inOrg = true
					}
				}
				if !inOrg {
					organizations = append(organizations, organization)
				}
			}

		}
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

func DeleteOrg(c *fiber.Ctx) error {
	var body struct {
		OrgId primitive.ObjectID `json:orgId`
	}
	c.BodyParser(&body)

	result, err := config.Database.Collection("organization").DeleteOne(context.TODO(), bson.D{
		{"_id", body.OrgId},
	})
	if err != nil {
		return c.JSON(fiber.Map{
			"status": 400,
			"error":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":   result,
		"status": 200,
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
			{"$set", bson.D{
				{"name", body.Organization.Name},
				{"location", body.Organization.Location},
				{"telephone", body.Organization.Telephone},
				{"tags", body.Organization.Tags},
				{"igurrl", body.Organization.IgUrl},
				{"description", body.Organization.Description},
				{"email", body.Organization.Email},
			}},
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

func SendMail(c *fiber.Ctx) error {
	var body models.Organization

	c.BodyParser(&body)

	m := gomail.NewMessage()

	adminMail := os.Getenv("ADMINMAIL")
	recieverMail := os.Getenv("RECIEVERMAIL")
	adminPass := os.Getenv("ADMINPASS")

	// Set E-Mail sender
	m.SetHeader("From", adminMail)

	// Set E-Mail receivers
	m.SetHeader("To", recieverMail)

	// Set E-Mail subject
	m.SetHeader("Subject", "nueva organizacion")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/html", `<h1>new organization</h1>
        <p>una nueva organizacion ah solicitado  registrarse, estos son sus datos</p>
        <ul>
            <li><b>name: </b>`+body.Name+`</li>
            <li><b>description: </b>`+body.Description+`</li>
            <li><b>telephone: </b>`+body.Telephone+`</li>
            <li><b>location: </b>`+body.Location+`</li>
            <li><b>email: </b>`+body.Email+`</li>
            <li><b>igtag: </b>`+body.Igtag+`</li>
        </ul>`)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, adminMail, adminPass)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return c.JSON(fiber.Map{
		"status": 200,
		"data":   "email sent to admin",
	})
}
