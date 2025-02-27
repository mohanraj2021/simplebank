package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	db       = "postgres"
	dbschema = "postgresql://root:admin@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

var testDb *sql.DB

func TestMain(m *testing.M) {
	var cerr error
	testDb, cerr = sql.Open(db, dbschema)

	if cerr != nil {
		log.Fatalf("Unable make connection to the DB beacaus %s", cerr.Error())
	}

	// defer conn.Close()

	testQueries = New(testDb)

	os.Exit(m.Run())

}
