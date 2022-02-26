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
	adminAuthenticated := admin.Use(middlewares.VerifyToken)
	adminAuthenticated.Get("about",controllers.Admin)
	adminAuthenticated.Put("email", controllers.UpdateAdmin)
	adminAuthenticated.Put("password", controllers.UpdatePassword)
	adminAuthenticated.Get("users",controllers.GetAllUser)
	//FEEDBACK
	adminAuthenticated.Get("feedback",controllers.GetAllFeedback)
	adminAuthenticated.Get("feedback/:id",controllers.GetOneFeedback)
	adminAuthenticated.Delete("feedback/:id",controllers.DeleteFeedback)
	//END FEEDBACK

	//GALLERY
	adminAuthenticated.Get("gallery", controllers.Gallery)
	adminAuthenticated.Get("gallery/:id", controllers.GetPhoto)
	adminAuthenticated.Put("gallery/:id", controllers.UpdatePhoto)
	adminAuthenticated.Delete("gallery/:id", controllers.DeletePhoto)
	adminAuthenticated.Post("gallery", controllers.CreatePhoto)
	//END GALLERY

	user := api.Group("user")
	user.Post("register", controllers.Register)
	user.Post("login",controllers.Login)
	//GALLERY
	user.Get("gallery", controllers.Gallery)
	user.Get("gallery/:id", controllers.GetPhoto)
	//END GALLERY
	userAuthenticated := user.Use(middlewares.VerifyToken)
	userAuthenticated.Get("about",controllers.User)

	feedback := api.Group("feedback")
	feedback.Post("", controllers.CreateFeedback)
}
