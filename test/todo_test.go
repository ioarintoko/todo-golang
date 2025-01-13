package test

import (
	"fmt"
	"testing"
	"time"
	"todolist/lib"
	"todolist/model"
)

var databrg = []*model.Todo{
	&model.Todo{IDTask: 1, Description: "Design Tampilan Mobile", Status: 1, CreateDate: time.Date(2018, time.March, 12, 0, 0, 0, 0, time.UTC), DueDate: time.Date(2018, time.March, 12, 0, 0, 0, 0, time.UTC), ExpireDate: time.Date(2018, time.March, 12, 0, 0, 0, 0, time.UTC)},
	&model.Todo{IDTask: 2, Description: "Meeting", Status: 1, CreateDate: time.Date(2018, time.March, 12, 0, 0, 0, 0, time.UTC), DueDate: time.Date(2018, time.March, 12, 0, 0, 0, 0, time.UTC), ExpireDate: time.Date(2018, time.March, 12, 0, 0, 0, 0, time.UTC)},
	&model.Todo{IDTask: 3, Description: "Futsal", Status: 1, CreateDate: time.Date(2018, time.March, 12, 0, 0, 0, 0, time.UTC), DueDate: time.Date(2018, time.March, 12, 0, 0, 0, 0, time.UTC), ExpireDate: time.Date(2018, time.March, 12, 0, 0, 0, 0, time.UTC)},
}

func TestTodo(t *testing.T) {
	t.Run("Test Insert Todo", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		defer db.Close()
		if err != nil {
			t.Fatal(err)
		}

		for _, val := range databrg {
			err := val.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
		}
	})
	t.Run("Test Update", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		defer db.Close()
		if err != nil {
			t.Fatal(err)
		}

		dataUpdate := map[string]interface{}{
			"description": "Bel Barang",
		}
		err = databrg[0].Update(db, dataUpdate)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Delete", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		defer db.Close()
		if err != nil {
			t.Fatal(err)
		}
		databrg := model.Todo{IDTask: 2}
		err = databrg.Delete(db)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Get", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		defer db.Close()
		if err != nil {
			t.Fatal(err)
		}
		databrg := model.Todo{IDTask: 1}
		err = databrg.Get(db)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(databrg)
	})

	t.Run("Test Gets", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		defer db.Close()
		if err != nil {
			t.Fatal(err)
		}
		databrg, err := model.GetsTodo(db)
		if err != nil {
			t.Fatal(err)
		}
		for _, val := range databrg {
			fmt.Println(*val)
		}
	})
}
