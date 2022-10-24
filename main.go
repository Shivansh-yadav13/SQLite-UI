package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shivansh-yadav13/sqlite-web/configs"
	"github.com/shivansh-yadav13/sqlite-web/db"
	"github.com/shivansh-yadav13/sqlite-web/routes"
)

func main() {
	app := configs.SetupApp()
	dbName := os.Getenv("SQLITE_NAME")
	if dbName == "" {
		dbName = "data"
	}
	db.ConnectDB(fmt.Sprintf("./sqlite_database/%s.db", dbName))
	routes.SetRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
