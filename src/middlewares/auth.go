package middlewares

import (
	"strconv"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func VerifyToken (c *fiber.Ctx) error{
	reqToken := c.Request().Header.Peek("Authorization")
	token, err := jwt.Parse(string(reqToken), func(t *jwt.Token) (interface{}, error) {
		return []byte("Umyt"), nil
	})
	if err != nil || !token.Valid{
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message":"Auth error",
		})
	}
	return c.Next()
}

func GetUser(c *fiber.Ctx) (uint, error) {
	reqToken := c.Request().Header.Peek("Authorization")

	token, err := jwt.ParseWithClaims(string(reqToken), &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("Umyt"),nil
	})

	if err != nil{
		return 0, err
	}

	payload := token.Claims.(*jwt.StandardClaims)

	id, _ := strconv.Atoi(payload.Subject)

	return uint(id), nil
}