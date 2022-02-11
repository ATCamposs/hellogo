package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		c.Response().Header.Add("Content-Type", "application/json")
		return c.JSON("Testando uso de memoria heroku")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"

	}
	log.Printf("Listening on port %s\n\n", port)
	app.Listen(":" + port)

}
