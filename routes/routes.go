package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	app.Get("/", index)
}

func index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "SQLite Web",
	})
}
