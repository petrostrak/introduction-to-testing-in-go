package main

import (
	"os"
	"simple-web-app/pkg/repository/dbrepo"
	"testing"
)

var app application

func TestMain(m *testing.M) {
	app.DB = &dbrepo.TestDBRepo{}
	app.Domain = "example.com"
	app.JWTSecret = "verysecret"

	os.Exit(m.Run())
}
