package router

import "github.com/gofiber/fiber/v2"

func RouteList(app *fiber.App) {
	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c *fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})
}
