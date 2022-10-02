package configs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

/*
SetupApp create a new fiber app
setting up the HTML templates and static files.
*/
func SetupApp() *fiber.App {
	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/static", "./static")
	return app
}
