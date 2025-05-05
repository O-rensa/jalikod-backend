package appservices

import (
	"context"

	irepositories "github.com/o-rensa/jalikod-backend/bffe/internal/domain/repositories"
)

type AuthenticationService struct {
	authenticationRepository irepositories.IAuthenticationRepository
}

func NewAuthenticationService(authenticationRepository irepositories.IAuthenticationRepository) *AuthenticationService {
	as := &AuthenticationService{}
	as.authenticationRepository = authenticationRepository
	return as
}

func (as *AuthenticationService) CheckUsername(ctx context.Context, username string) (string, error) {
	return as.authenticationRepository.GetUserUsername(ctx, username)
}

func (as *AuthenticationService) RegisterUser(ctx context.Context) (*RegisterResponse, error) {
	return nil, nil
}
