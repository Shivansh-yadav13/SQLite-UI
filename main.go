package main

import (
	"log"
	"shivansh-yadav13/sqlite-web/configs"
	"shivansh-yadav13/sqlite-web/routes"
)

func main() {
	app := configs.SetupApp()
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
