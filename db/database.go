package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() {
	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		fmt.Println(err)
	}

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	if err != nil {
		fmt.Println(err)
	}
	statement.Exec()
}
