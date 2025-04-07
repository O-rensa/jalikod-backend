package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/o-rensa/jalikod-backend/bffe/cmd/api"
	"github.com/o-rensa/jalikod-backend/bffe/initializers"
)

var db *sql.DB
var dbErr error

func init() {
	initializers.LoadDotEnv()
	db, dbErr = initializers.InitializePostgres()
	if dbErr != nil {
		log.Fatal(dbErr)
	}
}

func main() {
	server := api.NewAPIServer(os.Getenv("BFFE_PORT"), db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
