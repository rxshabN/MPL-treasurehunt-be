package main

import (
	"log"

	"github.com/Oik17/MPL-treasurehunt-be/internal/controllers"
	"github.com/Oik17/MPL-treasurehunt-be/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "pong",
		})
	})

	app.Post("/check-answer", controllers.CheckAnswer)

	log.Fatal(app.Listen(":3000"))
}
