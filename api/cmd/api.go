package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(map[string]any{
			"test": "ok",
		})
	})

	app.Listen(":8080")
}
