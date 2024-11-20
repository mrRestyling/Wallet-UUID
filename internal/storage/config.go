package storage

import (
	"log"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db *sqlx.DB
}

//

func New(conn *sqlx.DB) *Storage {
	return &Storage{
		Db: conn,
	}
}

func ConnectDB() *sqlx.DB {

	infoDB := Config()

	db, err := sqlx.Connect("postgres", infoDB)
	SayError(err, "Error connecting to database")

	err = db.Ping()
	SayError(err, "Error pinging database")

	return db

}

func SayError(err error, message string) {
	if err != nil {
		log.Fatal(message, err)
	}

}

func Config() string {

	host := "host=" + os.Getenv("HOST_WALLET")
	user := "user=" + os.Getenv("POSTGRES_USER")
	password := "password=" + os.Getenv("POSTGRES_PASSWORD")
	dbname := "dbname=" + os.Getenv("POSTGRES_DB")
	ssl := "sslmode=disable"

	qArr := []string{host, user, password, dbname, ssl}

	config := strings.Join(qArr, " ")

	return config

}
