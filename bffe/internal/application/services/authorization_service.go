package appservices

import (
	"context"

	irepositories "github.com/o-rensa/jalikod-backend/bffe/internal/domain/repositories"
)

type AuthorizationService struct {
	authorizationRepository irepositories.IAuthorizationRepository
}

func NewAuthorizationService(authorizationRepository irepositories.IAuthorizationRepository) *AuthorizationService {
	as := &AuthorizationService{}
	as.authorizationRepository = authorizationRepository
	return as
}

func (as *AuthorizationService) Hello(ctx context.Context) (string, error) {
	return as.authorizationRepository.GetUserUsername(ctx, "randomString")
}
