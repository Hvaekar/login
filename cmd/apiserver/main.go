package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "Port")

	srv := &http.Server{
		Addr: *addr,
		Handler: routes(),
	}

	log.Printf("Server started on localhost%s", *addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

