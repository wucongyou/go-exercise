package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int64
	Username string
	Password string
	Salt     string
	Tel      string
}

type Dao struct {
	db *sql.DB
}

func NewDAO(driverName, dataSourceName string) (res *Dao, err error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return
	}
	res = &Dao{
		db: db,
	}
	return
}

func main() {
	d, err := NewDAO("mysql", "testop:testop@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}

	u, err := d.UserByID(context.TODO(), 1)
	if err != nil {
		log.Print(err)
		return
	}

	str, _ := json.Marshal(u)
	fmt.Printf("res:%s", str)
}

func (d *Dao) UserByID(c context.Context, id int64) (res *User, err error) {
	rows := d.db.QueryRow("SELECT id,username,password,salt,tel FROM user WHERE id=?", id)
	res = new(User)
	if err = rows.Scan(&res.ID, &res.Username, &res.Password, &res.Salt, &res.Tel); err != nil {
		if err == sql.ErrNoRows {
			err = nil
			res = nil
			return
		}
	}
	return
}
