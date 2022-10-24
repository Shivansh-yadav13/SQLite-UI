package db

import (
	"fmt"
	"strings"
	"testing"
)

var dbName string = "test.db"

func TestConnectDB(t *testing.T) {
	if err := ConnectDB(dbName); err != nil {
		t.Error(err)
	}
}

func TestCreateTable(t *testing.T) {
	if err := ConnectDB(dbName); err != nil {
		t.Error(err)
	}
	newTableName := "userTable"
	if err := CreateTable(newTableName); err != nil {
		t.Fatal(err)
	}
	tables, err := GetTables()
	if err != nil {
		t.Fatal(err)
	}
	e := false
	for _, tb := range tables {
		if tb == newTableName {
			e = true
		}
	}
	if !e {
		t.Errorf("%s Table not found in the database.", newTableName)
	}
}

func TestDropTable(t *testing.T) {
	if err := ConnectDB(dbName); err != nil {
		t.Error(err)
	}
	dropTableName := "userTable"
	if err := DropTable(dropTableName); err != nil {
		t.Fatal(err)
	}
	tables, err := GetTables()
	if err != nil {
		t.Fatal(err)
	}
	for _, tb := range tables {
		if tb == dropTableName {
			t.Errorf("%s not deleted from the database.", dropTableName)
		}
	}
}

func TestGetSQLQuery(t *testing.T) {
	if err := ConnectDB(dbName); err != nil {
		t.Error(err)
	}
	newTableName := "tbl_schema_test"
	if err := CreateTable(newTableName); err != nil {
		t.Fatal(err)
	}
	sql, err := GetSQLQuery(newTableName)
	if err != nil {
		t.Fatal(err)
	}
	if sql == "" {
		t.Errorf("Table schema should not be empty")
	}
}

/*
	Test suite for AddColumn function
*/
func TestAddColumnForText(t *testing.T) {
	if err := ConnectDB(dbName); err != nil {
		t.Error(err)
	}
	newTableName := "tbl_schema_test"
	newColumnName := "demo_text"
	newColumnType := "TEXT"
	if err := AddColumn(newTableName, newColumnName, newColumnType); err != nil {
		t.Fatal(err)
	}

	sql := "SELECT sql FROM sqlite_schema where tbl_name=?"
	sql, err := GetSQLQuery(newTableName)
	if err != nil {
		t.Fatal(err)
	}
	if sql == "" {
		t.Errorf("Table schema should not be empty")
	}

	t.Log(sql)
	if !strings.Contains(sql, fmt.Sprintf("'%s' '%s'", newColumnName, newColumnType)) {
		t.Errorf("Column is not added correctly")
	}
}

func TestAddColumnForNull(t *testing.T) {
	if err := ConnectDB(dbName); err != nil {
		t.Error(err)
	}
	newTableName := "tbl_schema_test"
	newColumnName := "demo_null"
	newColumnType := "NULL"
	if err := AddColumn(newTableName, newColumnName, newColumnType); err != nil {
		t.Fatal(err)
	}

	sql := "SELECT sql FROM sqlite_schema where tbl_name=?"
	sql, err := GetSQLQuery(newTableName)
	if err != nil {
		t.Fatal(err)
	}
	if sql == "" {
		t.Errorf("Table schema should not be empty")
	}

	if !strings.Contains(sql, fmt.Sprintf("'%s' '%s'", newColumnName, newColumnType)) {
		t.Errorf("Column is not added correctly")
	}
}

func TestAddColumnForReal(t *testing.T) {
	if err := ConnectDB(dbName); err != nil {
		t.Error(err)
	}
	newTableName := "tbl_schema_test"
	newColumnName := "demo_real"
	newColumnType := "REAL"
	if err := AddColumn(newTableName, newColumnName, newColumnType); err != nil {
		t.Fatal(err)
	}

	sql := "SELECT sql FROM sqlite_schema where tbl_name=?"
	sql, err := GetSQLQuery(newTableName)
	if err != nil {
		t.Fatal(err)
	}
	if sql == "" {
		t.Errorf("Table schema should not be empty")
	}

	if !strings.Contains(sql, fmt.Sprintf("'%s' '%s'", newColumnName, newColumnType)) {
		t.Errorf("Column is not added correctly")
	}
}

func TestAddColumnForBlob(t *testing.T) {
	if err := ConnectDB(dbName); err != nil {
		t.Error(err)
	}
	newTableName := "tbl_schema_test"
	newColumnName := "demo_blob"
	newColumnType := "BLOB"
	if err := AddColumn(newTableName, newColumnName, newColumnType); err != nil {
		t.Fatal(err)
	}

	sql := "SELECT sql FROM sqlite_schema where tbl_name=?"
	sql, err := GetSQLQuery(newTableName)
	if err != nil {
		t.Fatal(err)
	}
	if sql == "" {
		t.Errorf("Table schema should not be empty")
	}

	if !strings.Contains(sql, fmt.Sprintf("'%s' '%s'", newColumnName, newColumnType)) {
		t.Errorf("Column is not added correctly")
	}
}
