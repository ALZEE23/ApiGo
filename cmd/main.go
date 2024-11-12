package main

import (
	"github.com/ALZEE23/ApiGo/database"
	"github.com/ALZEE23/ApiGo/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello,!")
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Test!")
	})

	app.Get("/fact", handlers.Home)

	app.Post("/fact", func(c *fiber.Ctx) error {
		return c.SendString("WTF ITS WORK?? YOLOOO")
	})

}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
