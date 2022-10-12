package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

/*
	ConnectDB connects the fiber app to the database.

	@param dbURI - database URI (eg: ./sqlite_database/data.db)
*/
func ConnectDB(dbUri string) error {
	var err error
	DB, err = sql.Open("sqlite3", dbUri)
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}

	return nil
}

/*
	GetTables gets list of tables from the database
*/
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

/*
	CreateTable adds a table to the database.

	@param tableName - name of the table
*/
func CreateTable(tableName string) error {
	_, err := DB.Exec(fmt.Sprintf("CREATE TABLE '%v'('id' INTEGER NOT NULL PRIMARY KEY)", tableName))
	if err != nil {
		return err
	}
	return nil
}

/*
	DeleteTable deletes a table from the database.

	@param tableName - name of the table
*/
func DropTable(tableName string) error {
	_, err := DB.Exec(fmt.Sprintf("DROP TABLE '%v'", tableName))
	if err != nil {
		return err
	}

	return nil
}

func GetSQLQuery(tableName string) (string, error) {
	sql := "SELECT sql FROM sqlite_master where tbl_name=?"
	params := []interface{}{tableName}
	stmt, err := DB.Prepare(sql)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	var sqltext string
	err = stmt.QueryRow(params...).Scan(&sqltext)
	if err != nil {
		return "", err
	}
	return sqltext, nil
}

/*
	AddColumn adds a column to a specified table.

	@param tableName - name of the table
	@param columnName - name of column to be added
	@param dataType - field type of the new column
*/
func AddColumn(tableName string, columnName string, dataType string) error {
	_, err := DB.Exec(fmt.Sprintf("ALTER TABLE '%s' ADD '%s' '%s'", tableName, columnName, dataType))
	if err != nil {
		return err
	}
	return nil
}