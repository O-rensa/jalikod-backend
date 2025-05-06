package irepositories

import (
	"context"

	domainaggregates "github.com/o-rensa/jalikod-backend/bffe/internal/domain/aggregates"
)

type IAuthenticationRepository interface {
	GetUserUsername(ctx context.Context, username string) (string, error)
	RegisterUser(ctx context.Context, usr domainaggregates.User, hashedpw string) (int32, error)
}
