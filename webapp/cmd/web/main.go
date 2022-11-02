package main

import (
	"encoding/gob"
	"flag"
	"log"
	"net/http"
	"simple-web-app/pkg/data"
	"simple-web-app/repository"
	"simple-web-app/repository/dbrepo"

	"github.com/alexedwards/scs/v2"
)

type application struct {
	DSN     string
	DB      repository.DatabaseRepo
	Session *scs.SessionManager
}

func main() {
	// register data.User{} with the session
	gob.Register(data.User{})

	// set up an app config
	app := application{
		Session: getSession(),
	}

	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5", "Postgres Connection")
	flag.Parse()

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}

	// get application routes
	mux := app.routes()

	// print out a msg
	log.Println("Starting server on port 8080...")

	// start the server
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
