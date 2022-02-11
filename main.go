package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		c.Response().Header.Add("Content-Type", "application/json")
		return c.JSON("Testando uso de memoria heroku")
	})

	app.Listen(":3000")
}
