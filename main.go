package main

import (
	"fiber/api/routes"
	"fiber/configs"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.DbConnect()
	app := fiber.New(fiber.Config{
		AppName:       "fiberkyy",
		ServerHeader:  "fiberkyy",
		StrictRouting: true,
		Prefork:       true,
		CaseSensitive: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	
	routes.UserRoute(app)

	app.Listen(":3000")
}
