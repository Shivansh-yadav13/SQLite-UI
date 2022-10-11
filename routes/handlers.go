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
		"Title":  "SQLite UI",
		"tables": tables,
	})
}

// Table will render the table_view html template.
func Table(c *fiber.Ctx) error {
	query, err := db.GetSQLQuery(c.Params("table"))
	if err != nil {
		return err
	}
	return c.Render("table_view", fiber.Map{
		"table_name": c.Params("table"),
		"sql_query":  query,
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
	table := new(TableStruct)
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
	table := new(TableStruct)
	if err := c.BodyParser(table); err != nil {
		log.Fatalln(err)
	}
	if err := db.DropTable(table.Name); err != nil {
		log.Fatalln(err)
	}

	return c.RedirectBack("/")
}

/*
	AddColumn takes the table name from the post request

	and calls the AddColumn of the db package.
*/
func AddColumn(c *fiber.Ctx) error {
	table := TableStruct{}
	column := ColumnStruct{}
	if err := c.BodyParser(&table); err != nil {
		return err
	}
	if err := c.BodyParser(&column); err != nil {
		return err
	}
	if err := db.AddColumn(table.Name, column.Name, column.Type); err != nil {
		log.Fatalln(err)
	}

	return c.RedirectBack(fmt.Sprintf("/%s", table.Name))
}
