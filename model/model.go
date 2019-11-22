package model

import "log"
import "database/sql"
import _ "github.com/mattn/go-sqlite3"

func Init() {
	log.Println("Init model")

	// connect to db
	db, err := sql.Open("sqlite3", "./database.db")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Println("Cannot ping database")
		log.Fatal(err)
	}

	// create tables
	_, err = db.Exec("create table if not exists test(id integer not null primary key);")
	if err != nil {
		log.Fatal(err)
	}

	// put some data
}
