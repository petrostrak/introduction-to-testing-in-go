package main

import (
	"log"
	"os"
	"simple-web-app/repository/dbrepo"
	"testing"
)

var app application

func TestMain(m *testing.M) {
	pathToTemlpates = "./../../templates/"
	app.Session = getSession()

	app.DSN = "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5"

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}

	os.Exit(m.Run())
}
