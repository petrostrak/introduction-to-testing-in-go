package main

import (
	"database/sql"
	"log"
)

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (app *application) connectToDB() (*sql.DB, error) {
	conn, err := openDB(app.DSN)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to Postgress!")

	return conn, nil
}
