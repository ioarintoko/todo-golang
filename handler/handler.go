package handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"todolist/model"
	// sessions "github.com/kataras/go-sessions"
)

var DB *sql.DB

func RegisDB(db *sql.DB) {

	DB = db
}

const (
	search           = "search"
	upload           = "upload"
	todo      		 = "todo"
	login            = "login"
)

func API(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")
	url := r.URL.Path
	dataURL := strings.Split(fmt.Sprintf("%v", url), "/")
	lastIndex := dataURL[len(dataURL)-1]
	dataParams := r.URL.Query().Get("params")
	switch dataURL[2] {
	case upload:
		switch r.Method {
		case http.MethodPost:
			if lastIndex == upload {
				var Buf bytes.Buffer
				// in your case file would be fileupload
				file, header, err := r.FormFile("file")
				if err != nil {
					panic(err)
				}
				fmt.Println(header.Filename)

				defer file.Close()
				name := strings.Split(header.Filename, ".")
				fmt.Printf("File name %s\n", name[0])
				img, err := os.Create(fmt.Sprintf("./front_end/image/%s", header.Filename))
				if err != nil {
					fmt.Println(err)
				}
				//img, err := os.Open(fmt.Sprintf("./image/%s", header.Filename))
				//byteContainer, err := ioutil.ReadAll(file) // why the long names though?
				//fmt.Printf("size:%d", len(byteContainer))
				// if err != nil {
				// 	fmt.Println(err)
				// }
				_, err = io.Copy(img, file)
				if err != nil {
					fmt.Println(err)
				}
				// Copy the file data to my buffer
				// do something with the contents...
				// I normally have a struct defined and unmarshal into a struct, but this will
				// work as an example
				//contents := Buf.String()
				//fmt.Println(contents)
				// I reset the buffer in case I want to use it again
				// reduces memory allocations in more intense projects
				Buf.Reset()
				// do something else
				// etc write header
				return
			}
		default:
			fmt.Println("Salah Method")
		}

	case todo:
		switch r.Method {
		case http.MethodGet:
			if lastIndex == todo || lastIndex == "" {
				databrg, err := model.GetsTodo(DB, dataParams)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}
				jsonData, err := json.Marshal(databrg)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}
				w.Write(jsonData)
			} else {
				id, err := strconv.Atoi(lastIndex)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				databrg := model.Todo{IDTask: id}
				err = databrg.Get(DB)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}
				jsonData, err := json.Marshal(databrg)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}
				w.Write(jsonData)
			}
		case http.MethodDelete:
			if lastIndex != todo {
				id, err := strconv.Atoi(lastIndex)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				databrg := model.Todo{IDTask: id}
				err = databrg.Delete(DB)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}
				w.Write([]byte("OK"))
			}
		case http.MethodPost:
			if lastIndex == todo {
				defer r.Body.Close()
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}
				//untuk mengubah json menjadi struct
				var Todo model.Todo
				err = json.Unmarshal(body, &Todo)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}
				//memasukkan data ke dalam database
				err = Todo.Insert(DB)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}
			}
		case http.MethodPut:
			if lastIndex != todo {
				url := r.URL
				dataURL := strings.Split(fmt.Sprintf("%v", url), "/")
				lastIndex := dataURL[len(dataURL)-1]

				defer r.Body.Close()
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}
				jsonMap := make(map[string]interface{})
				err = json.Unmarshal(body, &jsonMap)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}
				id, err := strconv.Atoi(lastIndex)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				databrg := model.Todo{IDTask: id}
				err = databrg.Update(DB, jsonMap)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}
			}
		default:
			fmt.Println("Salah Method")
		}
	}
}
