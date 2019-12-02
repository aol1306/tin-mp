package model

import (
	"database/sql"
	"encoding/hex"
	"log"
	"math/rand"
	"os"
	"time"

	"crypto/sha256"

	_ "github.com/mattn/go-sqlite3" // required for model
)

func openSQLConn() (*sql.DB, error) {
	// connect to db
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func hash(password string, salt string) string {
	sum := sha256.Sum256([]byte(password + salt))
	return hex.EncodeToString(sum[:])
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func getRandomSalt() string {
	return stringWithCharset(16, charset)
}

// AssingedUser represents a user assigned to card
type AssignedUser struct {
	id       int
	username string
}

// GetAssignedUsers gets users assigned to a card
func GetAssignedUsers(cardID int) []AssignedUser {
	db, err := openSQLConn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select user.id, user.username from user, card, user_card where user.id = user_card.id_user and card.id = user_card.id_card and user_card.id_card = ?;", cardID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	ret := []AssignedUser{}

	for rows.Next() {
		var id int
		var username string
		err := rows.Scan(&id, &username)
		if err != nil {
			log.Fatal(err)
		}
		ret = append(ret, AssignedUser{id, username})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return ret
}

// RegisterUser adds new user to db
func RegisterUser(username string, email string, password string) {
	randomSalt := getRandomSalt()
	db, err := openSQLConn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	passwordHash := hash(password, randomSalt)

	stmt, err := db.Prepare("insert into user(username, passwordHash, salt, email, active, admin, created) values (?, ?, ?, ?, 1, 0, datetime('now'))")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, passwordHash, randomSalt, email)
	if err != nil {
		log.Fatal(err)
	}
}

// VerifyUser checks if user login/pass is valid
func VerifyUser(username string, password string) bool {
	// connect to db
	db, err := openSQLConn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// get salt and hash for user
	var salt string
	var passwordhash string
	stmt, err := db.Prepare("select salt,passwordhash from user where username = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(username).Scan(&salt, &passwordhash)
	if err != nil {
		log.Println("No results - user or password invalid")
		return false
	}

	// hash password+salt
	currentHash := hash(password, salt)
	if currentHash == passwordhash {
		return true
	}
	return false
}

// AddCard adds new card to db
func AddCard(front string, back string, active int) {
	// front back created modified
	// connect to db
	db, err := openSQLConn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into card(front, back, active, created, modified) values(?, ?, ?, datetime('now'), datetime('now'));")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(front, back, active)
	if err != nil {
		log.Fatal(err)
	}
}

// DeleteCard removes card from db
func DeleteCard(id int) {
	// connect to db
	db, err := openSQLConn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("delete from card where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
}

// Card represents a card
type Card struct {
	ID     int
	Front  string
	Back   string
	Active int
}

// GetCardByID gets a card by ID
func GetCardByID(id int) []Card {
	// connect to db
	db, err := openSQLConn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select id,front,back,active from card where id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	ret := []Card{}
	for rows.Next() {
		var id int
		var front string
		var back string
		var active int
		err := rows.Scan(&id, &front, &back, &active)
		if err != nil {
			log.Fatal(err)
		}
		ret = append(ret, Card{id, front, back, active})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return ret
}

// GetAllCards returns all cards
func GetAllCards() []Card {
	// connect to db
	db, err := openSQLConn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select id, front, back, active from card")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	ret := []Card{}

	for rows.Next() {
		var id int
		var front string
		var back string
		var active int
		err := rows.Scan(&id, &front, &back, &active)
		if err != nil {
			log.Fatal(err)
		}
		ret = append(ret, Card{id, front, back, active})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return ret
}

// Init initializes the db and puts some default values
func Init() {
	os.Remove("database.db")

	log.Println("Init model")

	// connect to db
	db, err := openSQLConn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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
        passwordhash text not null,
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
        active integer default 1,
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
        srs_score integer default 0,
        last_seen text default '',
        count_wrong integer default 0,
        count_correct integer default 0,
        foreign key (id_user) references user (id)
        foreign key (id_card) references card (id)
    );
    `)
	if err != nil {
		log.Fatal(err)
	}

	// put some data

	// user:secret
	_, err = db.Exec(`
    insert into user(username, passwordhash, salt, email, active, admin, created)
	values ('user', 'f6ba18523c6942ba1e1b54f8256527ab1b8db94496cf6f4a2b6db9695c0fc6f9', 'abc', 'user@example.com', 1, 0, datetime('now'));
	insert into user(username, passwordhash, salt, email, active, admin, created)
	values ('user1', 'f6ba18523c6942ba1e1b54f8256527ab1b8db94496cf6f4a2b6db9695c0fc6f9', 'abc', 'user@example.com', 1, 0, datetime('now'));
	insert into user(username, passwordhash, salt, email, active, admin, created)
	values ('user2', 'f6ba18523c6942ba1e1b54f8256527ab1b8db94496cf6f4a2b6db9695c0fc6f9', 'abc', 'user@example.com', 1, 0, datetime('now'));
    `)
	if err != nil {
		log.Fatal(err)
	}

	// some cards
	_, err = db.Exec(`
	insert into card(front, back, active, created, modified) values ('人', 'człowiek', 1, datetime('now'), datetime('now'));
	insert into card(front, back, active, created, modified) values ('手', 'ręka', 1, datetime('now'), datetime('now'));
	insert into card(front, back, active, created, modified) values ('目', 'oko', 1, datetime('now'), datetime('now'));
	insert into card(front, back, active, created, modified) values ('天気', 'pogoda', 1, datetime('now'), datetime('now'));
	insert into card(front, back, active, created, modified) values ('年', 'rok', 1, datetime('now'), datetime('now'));
	insert into card(front, back, active, created, modified) values ('電車', 'pociąg', 1, datetime('now'), datetime('now'));
	insert into card(front, back, active, created, modified) values ('森', 'las', 1, datetime('now'), datetime('now'));
	insert into card(front, back, active, created, modified) values ('電池', 'bateria', 1, datetime('now'), datetime('now'));
	insert into card(front, back, active, created, modified) values ('結婚', 'ślub', 1, datetime('now'), datetime('now'));
	insert into card(front, back, active, created, modified) values ('車', 'samochód', 1, datetime('now'), datetime('now'));
	insert into card(front, back, active, created, modified) values ('勇者', 'bohater', 1, datetime('now'), datetime('now'));
	insert into card(front, back, active, created, modified) values ('感謝', 'wdzięczność', 1, datetime('now'), datetime('now'));
	insert into card(front, back, active, created, modified) values ('お金', 'pieniądze', 1, datetime('now'), datetime('now'));
	insert into card(front, back, active, created, modified) values ('気温', 'temperatura', 1, datetime('now'), datetime('now'));
	insert into card(front, back, active, created, modified) values ('着物', 'kimono', 1, datetime('now'), datetime('now'));
	insert into card(front, back, active, created, modified) values ('土曜日', 'Sobota', 1, datetime('now'), datetime('now'));
	`)
	if err != nil {
		log.Fatal(err)
	}

	// assign some cards to user 1
	_, err = db.Exec(`
	insert into user_card(id_user, id_card) values (1,1);
	insert into user_card(id_user, id_card) values (1,3);
	insert into user_card(id_user, id_card) values (1,4);
	insert into user_card(id_user, id_card) values (1,5);
	insert into user_card(id_user, id_card) values (1,8);
	`)
	if err != nil {
		log.Fatal(err)
	}
}
