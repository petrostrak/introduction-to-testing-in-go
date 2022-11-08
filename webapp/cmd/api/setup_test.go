package main

import (
	"os"
	"simple-web-app/pkg/repository/dbrepo"
	"testing"
)

var (
	app          application
	expiredToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiYXVkIjoiZXhhbXBsZS5jb20iLCJleHAiOjE2Njc1MzY3ODMsImlzcyI6ImV4YW1wbGUuY29tIiwibmFtZSI6IkpvaG4gRG9lIiwic3ViIjoiMSJ9.kiAYYa_vh9Gr8dqp-P3wg2969Rlwl9RmtxKC25Zpv6M"
)

func TestMain(m *testing.M) {
	app.DB = &dbrepo.TestDBRepo{}
	app.Domain = "example.com"
	app.JWTSecret = "verysecret"

	os.Exit(m.Run())
}
