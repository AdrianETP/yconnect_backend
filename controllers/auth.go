package controllers

import (
	"github.com/adrianetp/yconnect_backend/models"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var user models.UserLogin
	c.BodyParser(&user)
	if user.WebToken != "" {
		// login con web tokens
		return c.SendString("web token")
	} else {
		// login con usuarios
		return c.SendString("user")
	}
}
