package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/simplebank/api"
	db "github.com/simplebank/db/sqlc"
	"github.com/simplebank/types"
)

func main() {
	if lerr := godotenv.Load(".env"); lerr != nil {
		log.Fatalf("Unable make connection to the DB beacaus %s", lerr.Error())
		os.Exit(1)
	}

	types.DbDrive = os.Getenv("DB_DRIVE")
	types.DbSchema = os.Getenv("DB_SCHEMA")
	types.Addr = os.Getenv("ADDR")
	types.SecreteKey = os.Getenv("SECRETKEY")

	key, ok := os.LookupEnv("TOKEN_DURATION")
	if !ok {
		log.Fatal("Unable to get env variable ")
		os.Exit(1)
	}

	types.Token_Duration, _ = strconv.Atoi(key)

	conn, cerr := sql.Open(types.DbDrive, types.DbSchema)

	if cerr != nil {
		log.Fatalf("Unable make connection to the DB beacaus %s", cerr.Error())
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	server.Start(types.Addr)
}
