package configs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func SetupApp() *fiber.App {
	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "../static")
	return app
}
