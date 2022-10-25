package main

import (
	"log"
	"os"
	"strings"

	"github.com/shivansh-yadav13/sqlite-web/configs"
	"github.com/shivansh-yadav13/sqlite-web/db"
	"github.com/shivansh-yadav13/sqlite-web/routes"
)

func main() {
	app := configs.SetupApp()
	dbName := strings.TrimSpace(os.Getenv("SQLITE_NAME"))
	if dbName == "" {
		dbName = "data.db"
	}
	db.ConnectDB(dbName)
	routes.SetRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
