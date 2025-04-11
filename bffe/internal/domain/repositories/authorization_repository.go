package irepositories

import "context"

type IAuthorizationRepository interface {
	GetUserUsername(ctx context.Context, username string) (string, error)
}
