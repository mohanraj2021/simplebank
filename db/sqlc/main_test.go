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
	dbschema = "postgresql://root:admin@localhost:5432/simplebank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	conn, cerr := sql.Open(db, dbschema)

	if cerr != nil {
		log.Fatalf("Unable make connection to the DB beacaus %s", cerr.Error())
	}

	defer conn.Close()

	testQueries = New(conn)

	os.Exit(m.Run())

}
