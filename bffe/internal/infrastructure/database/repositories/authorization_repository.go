package dbrepositories

import (
	"context"
	"database/sql"

	sqlc "github.com/o-rensa/jalikod-backend/bffe/internal/infrastructure/database/generated"
)

type AuthenticationDBRepository struct {
	db *sqlc.Queries
}

func NewAuthenticationDBRepository(db *sql.DB) *AuthenticationDBRepository {
	r := &AuthenticationDBRepository{
		db: sqlc.New((db)),
	}
	return r
}

func (adbr *AuthenticationDBRepository) GetUserUsername(ctx context.Context, username string) (string, error) {
	return adbr.db.GetUserUsername(ctx, username)
}
