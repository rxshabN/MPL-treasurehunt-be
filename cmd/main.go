package main

import (
	"log"
	"time"

	"github.com/Oik17/MPL-treasurehunt-be/internal/controllers"
	"github.com/Oik17/MPL-treasurehunt-be/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.Connect()
	app := fiber.New()

	app.Use(cors.New())

	app.Use(logger.New(logger.Config{
		Format:     `{"time":"${time}","remote_ip":"${ip}","host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${ua}","status":${status},"latency":${latency},"latency_human":"${latency_human}","bytes_in":${bytes_received},"bytes_out":${bytes_sent}}` + "\n",
		TimeFormat: time.RFC3339Nano, 
		TimeZone:   "Local",          
	}))

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "pong",
		})
	})

	app.Post("/checkAnswer", controllers.CheckAnswer)
	app.Post("/createQuestion", controllers.CreateQuestion)
	log.Fatal(app.Listen(":8080"))
}
