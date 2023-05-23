package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("mysql", "root:si1v3r@tcp(127.0.0.1:3306)/test1")
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = DB.Ping()
	if err != nil {
		return err
	}
	log.Print("success connected")
	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Print("DB connection closed")
	}
}
