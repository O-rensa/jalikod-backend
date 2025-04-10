package database

import (
	"database/sql"

	dbrepositories "github.com/o-rensa/jalikod-backend/bffe/internal/infrastructure/database/repositories"
)

type DBRepositories struct {
	AuthorizationDBRepository *dbrepositories.AuthorizationDBRepository
}

func NewDBRepositories(db *sql.DB) *DBRepositories {
	dbr := &DBRepositories{}
	dbr.AuthorizationDBRepository = dbrepositories.NewAuthorizationDBRepository(db)
	return dbr
}
