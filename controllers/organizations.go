package controllers

import (
	"context"

	"github.com/adrianetp/yconnect_backend/config"
	"github.com/adrianetp/yconnect_backend/models"
	"github.com/gofiber/fiber/v2"
)

func CreateOrganization(c *fiber.Ctx) error{
	var Org models.Organization
	c.BodyParser(Org)
    result , error:= config.Database.Collection("organization").InsertOne(context.TODO() , Org)

    if error != nil {
        c.JSON(fiber.Map{
            "status":400,
            "error":error,

        })
    }

    return c.JSON(fiber.Map{
        "status":200,
        "data":result,
    })
}
