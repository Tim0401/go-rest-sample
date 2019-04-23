package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"./database"
)

type Message struct {
	DB   *sql.DB
	User *database.User
}
type Ping struct {
	Message string `json:"message"`
}

/**
 * TOPのルーティング
 */
func (m *Message) Handler(w http.ResponseWriter, r *http.Request) {

	var res []byte
	var err error

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		var p Ping
		p.Message = "Hello World!!"
		res, err = json.Marshal(p)
	default:
		http.NotFound(w, r)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

/**
 * Userのルーティング
 */
func (m *Message) UserHandler(w http.ResponseWriter, r *http.Request) {

	var res []byte
	var err error

	if r.URL.Path != "/users" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		data, _ := database.UsersAll(m.DB)
		if data == nil {
			res = []byte("[]")
		} else {
			res, err = json.Marshal(data)
		}
		w.WriteHeader(http.StatusOK)
	case http.MethodPost:
		m.User = new(database.User)
		m.User.Name = r.FormValue("name")
		m.User.Email = r.FormValue("email")
		data, _ := m.User.Insert(m.DB)
		res, err = json.Marshal(data)
		w.WriteHeader(http.StatusCreated)
	default:
		http.NotFound(w, r)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

/**
 * UserIdのルーティング
 */
func (m *Message) UserIdHandler(w http.ResponseWriter, r *http.Request) {

	var res []byte
	var err error

	str := r.URL.Path
	rep := regexp.MustCompile(`\s*/\s*`)
	result := rep.Split(str, -1)
	if result[1] != "users" || len(result) < 3 {
		http.NotFound(w, r)
		return
	}

	id := result[2]
	id64, e := strconv.ParseInt(id, 10, 64)
	if e != nil{
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		data, _ := database.UserByID(m.DB, id)
		if data == nil {
			http.NotFound(w, r)
			return
		}
		res, err = json.Marshal(data)
		w.WriteHeader(http.StatusOK)
	case http.MethodPut:
		m.User = new(database.User)
		m.User.Id = id64
		m.User.Name = r.FormValue("name")
		m.User.Email = r.FormValue("email")
		data, _ := m.User.Update(m.DB)
		res, err = json.Marshal(data)
		w.WriteHeader(http.StatusCreated)
	case http.MethodDelete:
		m.User = new(database.User)
		m.User.Id = id64
		err = m.User.Delete(m.DB)
		res = []byte("")
		w.WriteHeader(http.StatusNoContent)

	default:
		http.NotFound(w, r)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
	var m Message

	db, err := sql.Open("postgres", "host=db port=5432 user=user1 password=user1 dbname=go sslmode=disable")
	defer db.Close()

	if err != nil {
		fmt.Println(err)
		log.Fatalf("fail to connect database: %s", err)
	}

	m.DB = db
	http.HandleFunc("/users/", m.UserIdHandler)
	http.HandleFunc("/users", m.UserHandler)
	http.HandleFunc("/", m.Handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
