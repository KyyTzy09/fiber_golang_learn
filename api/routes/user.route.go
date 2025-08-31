package routes

import (
	"fiber/api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	user := app.Group("/user")
	user.Get("/", controllers.GetUsers)
	user.Get("/:id", controllers.GetUserById)
	user.Post("/create", controllers.CreateUser)
	user.Delete("/delete", controllers.DeleteAllUsers)
	user.Delete("/delete/:id", controllers.DeleteUserById)
}
