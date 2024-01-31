package controllers

import (
	"context"
	"time"

	"github.com/adrianetp/yconnect_backend/config"
	"github.com/adrianetp/yconnect_backend/encryption"
	"github.com/adrianetp/yconnect_backend/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

var secretKey = []byte("")

// Pendiente: Generar web token por cada log in
func generateToken(userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func validateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return "", err
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["sub"].(string)
	return userID, nil
}

func Login(c *fiber.Ctx) error {
	var userBody models.UserLogin
	c.BodyParser(&userBody)

	res := config.Database.Collection("Users").FindOne(context.TODO(), bson.D{{"email", userBody.Email}})

	var user models.User
	// Verify
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
	// Generar Token
	token, _ := generateToken(user.Id.String())
	return c.JSON(fiber.Map{
		"status": 200,
		"user":   user,
		"token":  token,
	})

}
