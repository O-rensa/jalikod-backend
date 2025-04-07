package initializers

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func InitializePostgres() (*sql.DB, error) {
	pgdb, err := sql.Open("postgres", os.Getenv("PG_DNS"))
	return pgdb, err
}
