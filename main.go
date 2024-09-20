package main

import "github.com/gofiber/fiber/v2"

func main() {
	// Create a fiber app server
	app := fiber.New()

	// Add a route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Miko!")
	})

	// Start the app server
	app.Listen(":3000")
}