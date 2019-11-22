package model

import "log"
import "database/sql"
import _ "github.com/mattn/go-sqlite3"

func Init() {
	log.Println("Init model")
    db, err := sql.Open("sqlite3", "./temp.db")
    if err != nil {
        log.Fatal(err)
    }
    log.Println("db OK")

    _, err = db.Exec("create table test(id integer not null primary key);")
    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()
}
