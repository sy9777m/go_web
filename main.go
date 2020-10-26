package main

import (
	"fmt"
	"net/http"
)

func main() {
	r := &router{make(map[string]map[string]http.HandlerFunc)}

	r.HandlerFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome!")
	})

	r.HandlerFunc("GET", "/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "About this website!")
	})

	r.HandlerFunc("GET", "/users/:id", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "retrieve user")
	})

	r.HandlerFunc("GET", "/users/:user_id/addresses/:address_id", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "retrieve user's address")
	})

	r.HandlerFunc("POST", "/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "create user")
	})

	r.HandlerFunc("POST", "/users/:user_id/addressed", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "create user's address")
	})

	http.ListenAndServe(":8000", r)
}
