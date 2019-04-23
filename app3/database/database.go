package database

import (
	"database/sql"
	_ "encoding/json"
	"log"
	"strconv"
)

type User struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func UsersAll(db *sql.DB) ([]*User, error) {

	rows, err := db.Query(`select * from users`)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rows.Close()

	var ms []*User
	for rows.Next() {
		m := &User{}
		// Tutorial 1-1. ユーザー名を表示しよう
		if err := rows.Scan(&m.Id, &m.Name, &m.Email, &m.CreatedAt, &m.UpdatedAt); err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	if err := rows.Err(); err != nil {
		log.Print(err)
		return nil, err
	}

	return ms, nil
}

func UserByID(db *sql.DB, id string) (*User, error) {
	m := &User{}

	// Tutorial 1-1. ユーザー名を表示しよう
	if err := db.QueryRow(`select id, username, email, created_at, updated_at from users where id = $1`, id).Scan(&m.Id, &m.Name, &m.Email, &m.CreatedAt, &m.UpdatedAt); err != nil {
		return nil, err
	}

	return m, nil
}

func (u *User) Insert(db *sql.DB) (*User, error) {

	var id int64
	err := db.QueryRow(`insert into users (username, email) values ($1,$2) RETURNING id`, u.Name, u.Email).Scan(&id)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	// 取得
	m, _ := UserByID(db, strconv.FormatInt(id, 10))
	return m, nil
}
func (u *User) Update(db *sql.DB) (*User, error) {
	_, err := db.Exec(`update users set username = $1 , email = $2, updated_at = now() where id = $3`, u.Name, u.Email, u.Id)

	if err != nil {
		log.Print(err)
		return nil, err
	}
	// 取得
	m, _ := UserByID(db, strconv.FormatInt(u.Id, 10))
	return m, nil
}
func (u *User) Delete(db *sql.DB) ( error) {
	_, err := db.Exec(`delete from users where id = $1`, u.Id)
	if err != nil {
		return err
	}
	return nil
}
