package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mysterybee07/jwt-auth/internal/controllers"
)

func Setup(app *fiber.App) {
	api := app.Group("/users")

	api.Get("/get-user", controllers.User)
	api.Post("/login", controllers.Login)
	api.Post("/logout", controllers.Logout)
	api.Post("/register", controllers.Register)
}
