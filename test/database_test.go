package test

import (
	"todolist/lib"
	"todolist/model"
	"testing"
)

var database, databaseDefaultMysql string

func init() {
	database = "todolist"
	databaseDefaultMysql = "Mysql"
}

func TestDatabaseMysql(t *testing.T) {
	t.Run("Testing Koneksi Mysql", func(t *testing.T) {
		db, err := lib.ConnectMySql(databaseDefaultMysql)

		defer db.Close()

		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Testing Drop", func(t *testing.T) {
		db, err := lib.ConnectMySql(databaseDefaultMysql)
		defer db.Close()
		if err != nil {
			t.Fatal(err)
		}
		err = lib.DropDB(db, database)
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Testing create", func(t *testing.T) {
		db, err := lib.ConnectMySql(databaseDefaultMysql)
		defer db.Close()
		if err != nil {
			t.Fatal(err)
		}
		err = lib.CreateDB(db, database)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Testing Create Table Todo", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		defer db.Close()
		if err != nil {
			t.Fatal(err)
		}
		err = lib.CreateTable(db, model.TableTodo)
		if err != nil {
			t.Fatal(err)
		}
	})

}