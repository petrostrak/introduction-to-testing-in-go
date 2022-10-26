package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

type contextKey string

const contextUserKey contextKey = "user_ip"

func (app *application) inFromContext(ctx context.Context) string {
	return ctx.Value(contextUserKey).(string)
}

func getIp(r *http.Request) (string, error) {
	// 192.0.0.0:8080
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "unknown", err
	}

	userIp := net.ParseIP(ip)
	if userIp == nil {
		return "", fmt.Errorf("userIp: %q is not IP:port", r.RemoteAddr)
	}

	// is it comming from a proxy?
	forward := r.Header.Get("X-Forwareded-For")
	if len(forward) > 0 {
		ip = forward
	}

	if len(ip) == 0 {
		ip = "forward"
	}

	return ip, nil
}

func (app *application) addIpToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		// get the ip
		ip, err := getIp(r)
		if err != nil {
			ip, _, _ = net.SplitHostPort(r.RemoteAddr)
			if len(ip) == 0 {
				ip = "unknown"
			}

			ctx = context.WithValue(r.Context(), contextUserKey, ip)
		} else {
			ctx = context.WithValue(r.Context(), contextUserKey, ip)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
