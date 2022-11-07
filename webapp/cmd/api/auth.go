package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	jwtTokenExpiry     = time.Minute * 15
	refreshTokenExpiry = time.Hour * 24
)

type TokenPairs struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	UserName string `json:"name"`
	jwt.RegisteredClaims
}

func (app *application) getTokenFromHeaderAndVerify(w http.ResponseWriter, r *http.Request) (string, *Claims, error) {
	//  we expect our authorization header to look like this:
	// Bearer <token>

	// add a header
	w.Header().Add("Vary", "Authorization")

	// get the authorization header
	authHeader := r.Header.Get("Authorization")

	// sanity check
	if authHeader == "" {
		return "", nil, errors.New("no auth header")
	}

	// split the header on spaces
	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return "", nil, errors.New("invalid auth header")
	}

	if headerParts[0] != "Bearer" {
		return "", nil, errors.New("unauthorized: no Bearer")
	}

	token := headerParts[1]

	// declare an empty Claims var
	claims := &Claims{}

	// parse the token with our claims (we read into claims), using our secret (from the receiver)
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		// validate the signing algorith
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(app.JWTSecret), nil
	})

	// check for an error; this catches expired tokens as well
	if err != nil {
		if strings.HasPrefix(err.Error(), "token is expired by") {
			return "", nil, errors.New("expired token")
		}

		return "", nil, err
	}

	// make sure that we issued this token
	if claims.Issuer != app.Domain {
		return "", nil, errors.New("incorrect issuer")
	}

	// valid token
	return token, claims, nil
}
