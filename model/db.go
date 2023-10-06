package model

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Setup() {
	// TODO: this should be changed - host is set to the IP of the container running the DB locally,
	dsn := `host=172.18.0.2 port=5432 user=admin password=test dbname=admin sslmode=disable`

	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Panic(err)
	}

	pingDB()
}

func pingDB() {
	err := db.Ping()
	if err != nil {
		log.Print(err)
	}
	log.Print("Successfully pinged database")
}
