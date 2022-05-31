package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
}

func GetTables() []string {
	rows, err := DB.Query("SELECT name FROM sqlite_schema WHERE type='table' ORDER BY name")
	if err != nil {
		panic(err)
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			panic(err)
		}
	}()

	tableNames := make([]string, 0)
	for rows.Next() {
		var tn string
		if err := rows.Scan(&tn); err != nil {
			panic(err)
		}
		tableNames = append(tableNames, tn)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	return tableNames
}

func CreateTable(tableName string) {
	rows, err := DB.Query(fmt.Sprintf("CREATE TABLE %d(ID INT PRIMARY KEY NOT NULL, )", tableName))
}
