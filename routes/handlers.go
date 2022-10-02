package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/shivansh-yadav13/sqlite-web/db"
)

// Index will render the index html template.
func Index(c *fiber.Ctx) error {
	tables, err := db.GetTables()
	if err != nil {
		panic(err)
	}
	return c.Render("index", fiber.Map{
		"Title":  "SQLite Web",
		"tables": tables,
	})
}

// GetTable will call the GetTables from the db package
func GetTables(c *fiber.Ctx) error {
	tNames, err := db.GetTables()
	var tNamesString string
	if err != nil {
		log.Fatalf("%v", err)
	}
	for i, name := range tNames {
		if i != len(tNames)-1 {
			tNamesString += name + ", "
			continue
		}
		tNamesString += name
	}
	return c.SendString(tNamesString)
}

/*
	CreateTable takes the table name from the post request

	and calls the CreateTable of the db package.
*/
func CreateTable(c *fiber.Ctx) error {
	table := new(Table)
	if err := c.BodyParser(table); err != nil {
		return err
	}
	if err := db.CreateTable(table.Name); err != nil {
		log.Fatalln(err)
	}

	return c.RedirectBack("/")
}

/*
	DropTable takes the table name from the post request

	and calls the DropTable of the db package.
*/
func DropTable(c *fiber.Ctx) error {
	table := new(Table)
	if err := c.BodyParser(table); err != nil {
		log.Fatalln(err)
	} else if err := db.DropTable(table.Name); err != nil {
		log.Fatalln(err)
	}
	return c.RedirectBack("/")
}