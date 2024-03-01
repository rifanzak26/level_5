package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Login(c *fiber.Ctx) error {
	user := c.FormValue("user")
	password := c.FormValue("password")

	if user != "admin" || password != "123456" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Making claim.
	claims := jwt.MapClaims{
		"name":  "admin",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Generate token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"token": t})
}