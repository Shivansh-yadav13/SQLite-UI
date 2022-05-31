package main

import (
	"github.com/shivansh-yadav13/sqlite-web/configs"
	"github.com/shivansh-yadav13/sqlite-web/db"
	"github.com/shivansh-yadav13/sqlite-web/routes"
	"log"
)

func main() {
	app := configs.SetupApp()
	db.ConnectDB()
	routes.SetRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
