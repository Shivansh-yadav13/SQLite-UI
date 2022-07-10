package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

var DB *sql.DB

func ConnectDB() error {
	var err error
	dbName := os.Getenv("SQLITE_NAME")
	if dbName == "" {
		dbName = "data"
	}
	DB, err = sql.Open("sqlite3", fmt.Sprintf("./sqlite_database/%s.db", dbName))
	if err != nil {
		return err
	}

	return nil
}

func GetTables() ([]string, error) {
	tableNames := make([]string, 0)
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

	return tableNames, nil
}

func CreateTable(tableName string) error {
	_, err := DB.Exec(fmt.Sprintf("CREATE TABLE %v('id' INTEGER NOT NULL PRIMARY KEY)", tableName))
	if err != nil {
		return err
	}
	return nil
}

func DropTable(tableName string) error {
	_, err := DB.Exec(fmt.Sprintf("DROP TABLE %v", tableName))
	if err != nil {
		return err
	}

	return nil
}
