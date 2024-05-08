package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/simplebank/api"
	db "github.com/simplebank/db/sqlc"
)

const (
	dbDrive  = "postgres"
	dbschema = "postgresql://root:admin@localhost:5432/simplebank?sslmode=disable"
	addr     = "0.0.0.0:2207"
)

func main() {
	conn, cerr := sql.Open(dbDrive, dbschema)

	if cerr != nil {
		log.Fatalf("Unable make connection to the DB beacaus %s", cerr.Error())
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	server.Start(addr)
}
