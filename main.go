package main

import (
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		c.Response().Header.Add("Content-Type", "application/json")
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		return c.JSON("Testando uso de memoria heroku: " + bToMb(m.Alloc) + "Mb")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"

	}
	log.Printf("Listening on port %s\n\n", port)
	app.Listen(":" + port)
}

func bToMb(b uint64) string {
	return strconv.FormatUint(b/1024/1024, 10)
}
