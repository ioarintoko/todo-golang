package model

import (
	"database/sql"
	"fmt"
	"strings"
	"todolist/lib"
	"time"
)

type Todo struct {
	IDTask    	 int    	`json:"idtask"`
	Description  string 	`json:"description"`
	CreateDate 	 time.Time 	`json:"createdate"`
	Status 		 int 		`json:"status"`
	DueDate 	 time.Time 	`json:"duedate"`
	ExpireDate 	 time.Time 	`json:"expiredate"`
}

var TableTodo = lib.Table{
	Name: "Todo",
	Field: []string{
		"IDTask INT(4) PRIMARY KEY  AUTO_INCREMENT",
		"Description VARCHAR(200)",
		"CreateDate TIMESTAMP",
		"Status INT(1)",
		"DueDate DATETIME",
		"ExpireDate DATETIME",
	},
}

func (m *Todo) Insert(db *sql.DB) error {
	query := "INSERT INTO Todo (Description, CreateDate, Status, DueDate, ExpireDate) VALUES (?,?,?,?,?)"
	_, err := db.Exec(query, m.Description, m.CreateDate, m.Status, m.DueDate, m.ExpireDate)
	return err
}

func (m *Todo) Delete(db *sql.DB) error {
	query := "DELETE FROM Todo WHERE IDTask = ?"
	_, err := db.Exec(query, m.IDTask)
	return err
}

func (m *Todo) Update(db *sql.DB, dataadm map[string]interface{}) error {
	var kolom = []string{}
	var args []interface{}
	for key, value := range dataadm {
		if value == "" {
			continue
		}
		updateData := fmt.Sprintf("%v = ?", strings.ToLower(key))
		kolom = append(kolom, updateData)
		args = append(args, value)
	}
	dataUpdate := strings.Join(kolom, ",")
	query := fmt.Sprintf("UPDATE Todo SET %s WHERE IDTask = '%d'", dataUpdate, m.IDTask)
	_, err := db.Exec(query, args...)
	fmt.Println(query)
	return err
}

func (m *Todo) Get(db *sql.DB) error {
	query := "SELECT * FROM Todo WHERE IDTask=?"
	err := db.QueryRow(query, &m.IDTask).Scan(&m.IDTask, &m.Description, &m.CreateDate, &m.Status, &m.DueDate, &m.ExpireDate)
	return err
}

func GetsTodo(db *sql.DB, params ...string) ([]*Todo, error) {
	var kolom = []string{}
	var args []interface{}
	if len(params) != 0 {
		if params[0] != "" {
			dataParams := strings.Split(params[len(params)-1], ";")
			for _, v := range dataParams {
				temp := strings.Split(fmt.Sprintf("%s", v), ",")
				where := fmt.Sprintf("%s %s ?", strings.ToLower(temp[0]), temp[1])
				kolom = append(kolom, where)
				//arg, _ := url.QueryUnescape(temp[2])
				args = append(args, temp[2])
			}
		}
	}
	dataKondisi := strings.Join(kolom, " AND ")
	var query string
	query = "SELECT * FROM Todo"
	if dataKondisi != "" {
		query += " WHERE " + dataKondisi
	}
	dataadm, err := db.Query(query, args...)

	if err != nil {
		return nil, err
	}
	defer dataadm.Close()
	var result []*Todo
	for dataadm.Next() {
		each := &Todo{}
		err := dataadm.Scan(&each.IDTask, &each.Description, &each.CreateDate, &each.Status, &each.DueDate, &each.ExpireDate)
		if err != nil {
			return nil, err
		}
		result = append(result, each)
	}
	return result, nil
}