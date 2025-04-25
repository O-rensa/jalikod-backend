package database

import (
	"database/sql"

	dbrepositories "github.com/o-rensa/jalikod-backend/bffe/internal/infrastructure/database/repositories"
)

type DBRepositories struct {
	AuthenticationDBRepository *dbrepositories.AuthenticationDBRepository
}

func NewDBRepositories(db *sql.DB) *DBRepositories {
	dbr := &DBRepositories{}
	dbr.AuthenticationDBRepository = dbrepositories.NewAuthenticationDBRepository(db)
	return dbr
}
