package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://nmea:qwerty@localhost:5433/nmea?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to the DB: %v", err)
	}

	informer := NewPGInformer(db)
	handler := NewPGConversationHandler(informer)
	// Listen on TCP port 2000 on all interfaces.
	nmeaServer := NewNmeaServer(":3000")

	nmeaServer.Serve(handler)
}
