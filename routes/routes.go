package routes

import (
	"github.com/gofiber/fiber/v2"
)

/*
	SetRoutes sets all the routes for the fiber app.
*/
func SetRoutes(app *fiber.App) {
	// template routes

	app.Get("/", Index)
	app.Get("/:table/schema", Table)
	app.Get("/:table/data", TableData)
	app.Get("/setting", Setting)

	// database ops routes

	app.Get("/get-tables", GetTables)
	app.Post("/create-table", CreateTable)
	app.Post("/drop-table", DropTable)
	app.Post("/add-column", AddColumn)
	// app.Post("add")
}
