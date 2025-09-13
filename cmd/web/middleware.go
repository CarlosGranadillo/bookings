package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// WriteToConsoleOnPageViewed writes to the console when a page is viewed
func WriteToConsoleOnPageViewed(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Page was viewed")
			next.ServeHTTP(w, r)
		})
}

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   appConfig.IsProd,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler

}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
