package appservices

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v3"
	domainaggregates "github.com/o-rensa/jalikod-backend/bffe/internal/domain/aggregates"
	irepositories "github.com/o-rensa/jalikod-backend/bffe/internal/domain/repositories"
	domainutilities "github.com/o-rensa/jalikod-backend/bffe/internal/domain/utility"
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

func (as *AuthenticationService) RegisterUser(ctx context.Context, req RegisterRequest) (int, RegisterResponse, error) {
	mi := req.MiddleInitial
	ext := req.NameExtension
	hashedpw, err := domainutilities.HashPassword(req.Password)
	if err != nil {
		em := "fail to hash password"
		res := RegisterResponse{}
		res.Status = RegisterFail
		res.FailMessage = append(res.FailMessage, Other)
		res.ErrorMessage = &em
		return fiber.StatusInternalServerError, res, err
	}
	usr := domainaggregates.NewUser(0, req.Username)
	usr.SetFullName(req.FirstName, *mi, req.Surname, *ext)
	usr.CreationTime = time.Now()
	usr.CreatorUserid = 0
	if _, err := as.authenticationRepository.RegisterUser(ctx, *usr, hashedpw); err != nil {
		res := RegisterResponse{}
		res.Status = RegisterFail
		return fiber.StatusInternalServerError, res, err
	}
	res := RegisterResponse{}
	res.Status = RegisterSuccess
	return fiber.StatusCreated, res, nil
}
