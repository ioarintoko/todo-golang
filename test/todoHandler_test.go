package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"todolist/handler"
	"todolist/lib"
	"todolist/model"
	"testing"
	"time"
)

func TestHandlerTodo(t *testing.T) {
	t.Run("Testing Insert", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		defer db.Close()
		if err != nil {
			t.Fatal(err)
		}
		dataTodo := []model.Todo{
			{IDTask: 1, Description: "Design Tampilan Mobile", Status: 1, CreateDate: time.Date(2018, time.March, 12, 0, 0, 0, 0, time.UTC), DueDate: time.Date(2018, time.March, 12, 0, 0, 0, 0, time.UTC), ExpireDate: time.Date(2018, time.March, 12, 0, 0, 0, 0, time.UTC)},
			{IDTask: 2, Description: "Meeting", Status: 1, CreateDate: time.Date(2018, time.March, 12, 0, 0, 0, 0, time.UTC), DueDate: time.Date(2018, time.March, 12, 0, 0, 0, 0, time.UTC), ExpireDate: time.Date(2018, time.March, 12, 0, 0, 0, 0, time.UTC)},
			{IDTask: 3, Description: "Futsal", Status: 1, CreateDate: time.Date(2018, time.March, 12, 0, 0, 0, 0, time.UTC), DueDate: time.Date(2018, time.March, 12, 0, 0, 0, 0, time.UTC), ExpireDate: time.Date(2018, time.March, 12, 0, 0, 0, 0, time.UTC)},
		}

		for _, val := range dataTodo {
			jsonData, err := json.Marshal(val)
			if err != nil {
				t.Fatal(err)
			}
			b := bytes.NewBuffer(jsonData)
			r := httptest.NewRequest(http.MethodPost, "/api/todo", b)
			w := httptest.NewRecorder()
			handler.RegisDB(db)
			handler.API(w, r)

			status := w.Code
			if status != http.StatusOK {
				t.Fatalf("get : %v want : %v status : %v", status, http.StatusOK, w.Body)
			}
		}
	})

	t.Run("Testing Update", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		defer db.Close()
		if err != nil {
			t.Fatal(err)
		}
		dataTodo := model.Todo{
			Description: "Lari Pagi",
		}
		jsonData, err := json.Marshal(dataTodo)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(jsonData)
		fmt.Println(dataTodo)
		b := bytes.NewBuffer(jsonData)
		r := httptest.NewRequest(http.MethodPost, "/api/todo/1", b)
		w := httptest.NewRecorder()
		handler.RegisDB(db)
		handler.API(w, r)

		status := w.Code
		if status != http.StatusOK {
			t.Fatalf("get : %v want : %v status : %v", status, http.StatusOK, w.Body)
		}
	})

	t.Run("Testing Delete", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		defer db.Close()
		if err != nil {
			t.Fatal(err)
		}
		r := httptest.NewRequest(http.MethodDelete, "/api/todo/2", nil)
		w := httptest.NewRecorder()
		handler.RegisDB(db)
		handler.API(w, r)

		status := w.Code
		if status != http.StatusOK {
			t.Fatalf("get : %v want : %v status : %v", status, http.StatusOK, w.Body)
		}
	})

	t.Run("Testing Gets", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		defer db.Close()
		if err != nil {
			t.Fatal(err)
		}
		r := httptest.NewRequest(http.MethodGet, "/api/todo", nil)
		w := httptest.NewRecorder()
		handler.RegisDB(db)
		handler.API(w, r)

		status := w.Code
		if status != http.StatusOK {
			t.Fatalf("get : %v want : %v status : %v", status, http.StatusOK, w.Body)
		}

		var jsonData = []byte(w.Body.String())
		var todo []model.Todo

		err = json.Unmarshal(jsonData, &todo)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(todo)
	})

	//getsWithParams
	t.Run("Testing GetsWithParams", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		defer db.Close()
		if err != nil {
			t.Fatal(err)
		}
		r := httptest.NewRequest(http.MethodGet, "/api/todo?params=idtask%2Clike%2C%251%25", nil)
		w := httptest.NewRecorder()
		handler.RegisDB(db)
		handler.API(w, r)

		status := w.Code
		if status != http.StatusOK {
			t.Fatalf("get : %v want : %v status : %v", status, http.StatusOK, w.Body)
		}

		var jsonData = []byte(w.Body.String())
		var todo []model.Todo

		err = json.Unmarshal(jsonData, &todo)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(todo)
	})

	t.Run("Testing Get", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		defer db.Close()
		if err != nil {
			t.Fatal(err)
		}
		r := httptest.NewRequest(http.MethodGet, "/api/todo/1", nil)
		w := httptest.NewRecorder()
		handler.RegisDB(db)
		handler.API(w, r)

		status := w.Code
		if status != http.StatusOK {
			t.Fatalf("get : %v want : %v status : %v", status, http.StatusOK, w.Body)
		}

		var jsonData = []byte(w.Body.String())
		var todo model.Todo

		err = json.Unmarshal(jsonData, &todo)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(todo)
	})
}
