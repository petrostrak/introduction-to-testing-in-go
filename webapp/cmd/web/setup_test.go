package main

import (
	"os"
	"simple-web-app/repository/dbrepo"
	"testing"
)

var app application

func TestMain(m *testing.M) {
	pathToTemlpates = "./../../templates/"
	app.Session = getSession()

	app.DB = &dbrepo.TestDBRepo{}

	os.Exit(m.Run())
}
