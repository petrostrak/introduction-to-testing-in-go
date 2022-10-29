package main

import (
	"os"
	"testing"
)

var app application

func TestMain(m *testing.M) {
	pathToTemlpates = "./../../templates/"
	app.Session = getSession()

	os.Exit(m.Run())
}
