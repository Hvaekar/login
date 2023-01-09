package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
)

type app struct {
	DB *sql.DB
}

func main() {
	addr := flag.String("addr", ":4000", "Port")

	db, err := openDB("web:ukraine@/somedb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := &app{
		DB: db,
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	log.Printf("Server started on localhost%s", *addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
