package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ryanadiputraa/spotwave/api/config"
)

func main() {
	config, err := config.LoadConfig("config")
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(map[string]any{
			"test": "ok",
		})
	})

	app.Listen(config.Port)
}
