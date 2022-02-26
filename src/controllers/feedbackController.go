package controllers

import (
	"memorise/src/database"
	"memorise/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)
func GetAllFeedback(c *fiber.Ctx) error {
	var feedbacks []models.Feedback
	database.DB.Find(&feedbacks)
	return c.JSON(feedbacks)
}
func GetOneFeedback(c *fiber.Ctx) error{
	id, _ := strconv.Atoi(c.Params("id"))
	var feedback models.Feedback
	feedback.Id = uint(id)
	database.DB.Find(&feedback)
	return c.JSON(feedback)
}
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
func DeleteFeedback(c *fiber.Ctx) error{
	id, _ := strconv.Atoi(c.Params("id"))
	feedback := models.Feedback{
		Id: uint(id),
	}
	database.DB.Delete(&feedback)
	return c.JSON(fiber.Map{
		"message":"success",
	})
}