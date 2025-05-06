package dbrepositories

import (
	"context"
	"database/sql"

	domainaggregates "github.com/o-rensa/jalikod-backend/bffe/internal/domain/aggregates"
	domainutilities "github.com/o-rensa/jalikod-backend/bffe/internal/domain/utility"
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

func (adbr *AuthenticationDBRepository) RegisterUser(ctx context.Context, usr domainaggregates.User, hashedpw string) (int32, error) {
	param := sqlc.RegisterUserAndGetIdParams{}
	fn, mn, sn, ext := usr.GetFullname()
	param.FirstName = fn
	param.MiddleInitial = domainutilities.ToNullString(mn)
	param.Surname = sn
	param.NameExtension = domainutilities.ToNullString(ext)
	return adbr.db.RegisterUserAndGetId(ctx, param)
}
