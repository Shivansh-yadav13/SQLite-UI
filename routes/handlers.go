package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shivansh-yadav13/sqlite-web/db"
)

func Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "SQLite Web",
	})
}

func GetTables(c *fiber.Ctx) error {
	tNames := db.GetTables()
	return c.SendString(tNames[0])
}
