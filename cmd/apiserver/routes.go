package main

import "net/http"

func (app *app) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/login", app.login)

	return mux
}
