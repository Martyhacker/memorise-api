package controllers

import (
	"memorise/src/database"
	"memorise/src/models"

	"github.com/gofiber/fiber/v2"
)

func CreateFeedback(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil{
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"Error occured on creating feedback",
		})
	}

	feedback := models.Feedback{
		Email: data["email"],
		Phone: data["phone"],
		Content: data["content"],
	}
	database.DB.Create(&feedback)

	return c.JSON(feedback)
}