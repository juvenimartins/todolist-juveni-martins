package model

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConectDb() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "root:root@(localhost:3306)/todolist")
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
