package main

import (
	"PostgresCRUD/router"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	host     = "localhost"
	port     = 5555
	user     = "postgres"
	password = "postgres"
	dbname   = "test"
)

func main() {
	// create the postgres db connection
	db := MigrateDB()

	// close the db connection
	defer db.Close()

	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))

}

// Migrate DB set up
func MigrateDB() *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open the connection
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	m, err := migrate.New(
		"file://db/migrations/initialize",
		"postgres://postgres:postgres@localhost:5555/test?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	m.Up()

	if err != nil {
		panic(err)
	}

	return db
}
