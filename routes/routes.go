package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {

	// template routes

	app.Get("/", Index)
	//app.Post("/create-table")

	// database ops routes

	app.Get("/get-tables", GetTables)
}
