package routes

import (
	"memorise/src/controllers"
	"memorise/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	admin := api.Group("admin")
	admin.Post("login",controllers.LoginAdmin)
	adminAuthenticated := admin.Use(middlewares.IsAuthenticated)
	adminAuthenticated.Get("about",controllers.Admin)
	adminAuthenticated.Post("logout",controllers.LogoutAdmin)
	adminAuthenticated.Put("email", controllers.UpdateAdmin)
	adminAuthenticated.Put("password", controllers.UpdatePassword)
	adminAuthenticated.Get("gallery", controllers.Gallery)
	adminAuthenticated.Get("gallery/:id", controllers.GetPhoto)
	adminAuthenticated.Put("gallery/:id", controllers.UpdatePhoto)
	adminAuthenticated.Delete("gallery/:id", controllers.DeletePhoto)
	adminAuthenticated.Post("gallery", controllers.CreatePhoto)

	user := api.Group("user")
	user.Post("register", controllers.Register)
	user.Post("login",controllers.Login)
	user.Get("gallery", controllers.Gallery)
	user.Get("gallery/:id", controllers.GetPhoto)
	userAuthenticated := user.Use(middlewares.IsAuthenticated)
	userAuthenticated.Get("about",controllers.User)
	userAuthenticated.Post("logout",controllers.Logout)


	

	feedback := api.Group("feedback")
	feedback.Post("", controllers.CreateFeedback)
}
