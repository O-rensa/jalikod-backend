package dbrepositories

import (
	"database/sql"

	sqlc "github.com/o-rensa/jalikod-backend/bffe/internal/infrastructure/database/generated"
)

type AuthorizationDBRepository struct {
	db *sqlc.Queries
}

func NewAuthorizationDBRepository(db *sql.DB) *AuthorizationDBRepository {
	r := &AuthorizationDBRepository{
		db: sqlc.New((db)),
	}
	return r
}

func (adbr *AuthorizationDBRepository) Test() string {
	return "HELLO FROM AUTHORIZATION REPOSITORY"
}
