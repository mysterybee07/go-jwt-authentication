package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mysterybee07/jwt-auth/internal/database"
	"github.com/mysterybee07/jwt-auth/internal/routes"
)

func init() {
	database.DBconn()
	database.LoadEnvironment()
}
func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:8080",                       // Specify the allowed origins
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",                 // Specify the allowed methods
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization", // Specify the allowed headers
	}))

	// app.Get("/", func(c fiber.Ctx) error {
	// 	return c.SendString("Hello World")
	// })
	routes.Setup(app)

	app.Listen(":8080")
}
