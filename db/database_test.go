package db

import (
	"testing"
)

func TestConnectDB(t *testing.T) {
	if err := ConnectDB("../sqlite_database/data.db"); err != nil {
		t.Error(err)
	}
}

func TestCreateTable(t *testing.T) {
	if err := ConnectDB("../sqlite_database/data.db"); err != nil {
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
	if err := ConnectDB("../sqlite_database/data.db"); err != nil {
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
	if err := ConnectDB("../sqlite_database/data.db"); err != nil {
		t.Error(err)
	}
	newTableName := "userTable"
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
