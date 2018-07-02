package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"github.com/gorilla/sessions"
	"math/rand"
)

// Db handle
var Db *sql.DB
var err error
var CookieStore *sessions.CookieStore

//func InitSQLiteDB() *sql.DB {
func InitSQLiteDB() {

	Db, err = sql.Open("sqlite3", "./config/borgdirmedia.db")
	test(err)

	key := make([]byte, 10)
	rand.Read(key)
	CookieStore = sessions.NewCookieStore(key)

	// gob.Register(&EquipInfo{})
	// gob.Register(&CartData{})
}

func ReturnDB() *sql.DB {
	return Db
}

func ReturnCS() *sessions.CookieStore {
	return CookieStore
}

func test(e error) {
	if (err != nil) {
		fmt.Println("FEHLER : ", e)
	}
}
