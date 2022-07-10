package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/shivansh-yadav13/sqlite-web/db"
)

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

func CreateTable(c *fiber.Ctx) error {
	table := new(Table)
	if err := c.BodyParser(table); err != nil {
		return err
	}
	if err := db.CreateTable(table.Name); err != nil {
		log.Fatalln(err)
	}

	return c.Redirect("/", 201)
}

func DropTable(c *fiber.Ctx) error {
	table := new(Table)
	if err := c.BodyParser(table); err != nil {
		log.Fatalln(err)
	}
	if err := db.DropTable(table.Name); err != nil {
		log.Fatalln(err)
	}

	return c.Redirect("/", 202)
}
