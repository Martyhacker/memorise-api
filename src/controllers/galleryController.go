package controllers

import (
	"fmt"
	"memorise/src/database"
	"memorise/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Gallery(c *fiber.Ctx) error {
	var photos []models.Gallery
	database.DB.Find(&photos)
	return c.JSON(photos)
}

func CreatePhoto(c *fiber.Ctx) error {
	file, err := c.FormFile("photo")

	if err != nil {
		return err
	}
	photo := models.Gallery{
		Photo: "./uploads/" + file.Filename,
	}
	database.DB.Create(&photo)
	c.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
	return c.JSON(photo)
}

func GetPhoto(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var photo models.Gallery

	photo.Id = uint(id)

	database.DB.Find(&photo)
	return c.JSON(photo)
}
func UpdatePhoto(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	file, err := c.FormFile("photo")
	
	photo := models.Gallery{
		Id: uint(id),
		Photo: "./uploads/" + file.Filename,
	}

	if err != nil {
		return err
	}
	
	database.DB.Model(&photo).Updates(&photo)
	return c.JSON(photo)
}

func DeletePhoto(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	photo := models.Gallery{
		Id: uint(id),
	}
	database.DB.Delete(&photo)
	return c.JSON(fiber.Map{
		"message":"success",
	})
}
