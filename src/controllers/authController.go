package controllers

import (
	"memorise/src/database"
	"memorise/src/middlewares"
	"memorise/src/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match",
		})
	}

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}
	user.SetPassword(data["password"])

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		database.DB.Create(&user)
		payload := jwt.StandardClaims{
			Subject:   strconv.Itoa(int(user.Id)),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		}

		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("Umyt"))
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": "Invalid Credentials",
			})
		}

		return c.JSON(fiber.Map{
			"message": "success",
			"token":   token,
		})
	} else {
		return c.JSON(fiber.Map{
			"message": "Already exists",
		})
	}

}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	payload := jwt.StandardClaims{
		Subject:   strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("Umyt"))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"token":token,
	})
}

func User(c *fiber.Ctx) error {
	id, _ := middlewares.GetUser(c)
	var user models.User
	database.DB.Where("id = ?", id).First(&user)

	return c.JSON(user)
}
