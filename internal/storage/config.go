package storage

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db *sqlx.DB
}

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

	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(1 * time.Minute)

	return db
}

func SayError(err error, message string) {
	if err != nil {
		log.Fatal(message, err)
	}

}

func Config() string {

	// host := "host=" + os.Getenv("HOST_WALLET") // локально
	host := "host=postgres" // docker
	user := "user=" + os.Getenv("POSTGRES_USER")
	password := "password=" + os.Getenv("POSTGRES_PASSWORD")
	dbname := "dbname=" + os.Getenv("POSTGRES_DB")
	ssl := "sslmode=disable"

	qArr := []string{host, user, dbname, ssl, password}

	config := strings.Join(qArr, " ")

	// time.Sleep(3 * time.Second) // для контейнеров

	return config

}
