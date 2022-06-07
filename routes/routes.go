package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {

	// template routes

	app.Get("/", Index)

	// database ops routes

	app.Get("/get-tables", GetTables)
	app.Post("/create-table", CreateTable)
	app.Post("/drop-table", DropTable)
	//app.Post("/add-col", AddCol)
	//app.Post("add")
}
