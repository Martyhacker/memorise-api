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

// func RegisterAdmin(c *fiber.Ctx) error {
// 	var data map[string]string

// 	if err := c.BodyParser(&data); err != nil {
// 		return err
// 	}
// 	if data["password"] != data["password_confirm"] {
// 		c.Status(400)
// 		return c.JSON(fiber.Map{
// 			"message": "Passwords do not match",
// 		})
// 	}

// 	admin := models.User{
// 		Email:        data["email"],
// 	}
// 	admin.SetPassword(data["password"])

// 	database.DB.Create(&admin)

// 	return c.JSON(admin)
// }

func LoginAdmin(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var admin models.Admin

	database.DB.Where("email = ?", data["email"]).First(&admin)

	if admin.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}
	if err := admin.ComparePassword(data["password"]); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	payload := jwt.StandardClaims{
		Subject:   strconv.Itoa(int(admin.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("Umyt"))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func Admin(c *fiber.Ctx) error{
	id, _ := middlewares.GetUser(c)
	var admin models.Admin
	database.DB.Where("id = ?", id).First(&admin)

	return c.JSON(admin)
}

func LogoutAdmin(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message":"success",
	})
}

func UpdateAdmin(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err!=nil{
		return err
	}
	id, _ := middlewares.GetUser(c)

	admin := models.Admin{
		Id: id,
		Email: data["email"],
	}
	database.DB.Model(&admin).Updates(&admin)

	return c.JSON(admin)
}
func UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err!=nil{
		return err
	}

	if data["password"] != data["password_confirm"]{
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"Passwords does't match",
		})
	}

	id, _ := middlewares.GetUser(c)

	admin := models.Admin{
		Id: id,
	}
	admin.SetPassword(data["password"])
	database.DB.Model(&admin).Updates(&admin)

	return c.JSON(admin)
}