package irepositories

import "context"

type IAuthenticationRepository interface {
	GetUserUsername(ctx context.Context, username string) (string, error)
}
