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
	log.Println("db ok")

	// create tables
	_, err = db.Exec(`
    create table if not exists user
    (
        id integer not null primary key,
        username text not null,
        password text not null,
        salt text not null,
        email text not null,
        active integer not null,
        admin integer not null,
        created text not null
    );
    `)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
    create table if not exists card
    (
        id integer not null primary key,
        front text,
        back text,
        active integer,
        created text,
        modified text
    );
    `)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
    create table if not exists user_card
    (
        id integer not null primary key,
        id_user integer not null,
        id_card integer not null,
        srs_score integer,
        last_seen text,
        count_wrong integer,
        count_correct integer,
        foreign key (id_user) references user (id)
        foreign key (id_card) references card (id)
    );
    `)
	if err != nil {
		log.Fatal(err)
	}

	// put some data
}
