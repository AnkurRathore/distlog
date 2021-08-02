package main

import (
	"log"

	"github.com/AnkurRathore/distlog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8000")
	log.Fatal(srv.ListenAndServe())
}
