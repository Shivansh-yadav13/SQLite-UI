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
	app.Get("/:table", Table)

	// database ops routes

	app.Get("/get-tables", GetTables)
	app.Post("/create-table", CreateTable)
	app.Post("/drop-table", DropTable)
	// app.Post("/add-col", AddCol)
	// app.Post("add")
}
