package dbrepositories

import (
	"context"
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

func (adbr *AuthorizationDBRepository) GetUserUsername(ctx context.Context, username string) (string, error) {
	return adbr.db.GetUserUsername(ctx, username)
}
