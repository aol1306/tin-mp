package model

import "log"
import "time"
import "encoding/hex"
import "database/sql"
import "math/rand"
import _ "github.com/mattn/go-sqlite3"
import "crypto/sha256"

func openSqlConn() (*sql.DB, error) {
	// connect to db
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return nil, err
	} else {
		return db, nil
	}
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

func RegisterUser(username string, email string, password string) {
	randomSalt := getRandomSalt()
	db, err := openSqlConn()
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

func VerifyUser(username string, password string) bool {
	// connect to db
	db, err := openSqlConn()
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
	} else {
		return false
	}
}

func Init() {
	log.Println("Init model")

	// connect to db
	db, err := openSqlConn()
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

	// user:secret
	_, err = db.Exec(`
    insert into user(username, passwordhash, salt, email, active, admin, created)
    values ('user', 'f6ba18523c6942ba1e1b54f8256527ab1b8db94496cf6f4a2b6db9695c0fc6f9', 'abc', 'user@example.com', 1, 0, datetime('now'));
    `)
	if err != nil {
		log.Fatal(err)
	}
}
