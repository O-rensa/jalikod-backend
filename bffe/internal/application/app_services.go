package application

import (
	appservices "github.com/o-rensa/jalikod-backend/bffe/internal/application/services"
	"github.com/o-rensa/jalikod-backend/bffe/internal/infrastructure/database"
)

type AppServices struct {
	AuthorizationService *appservices.AuthorizationService
}

func NewApplicationServices(dbRepositories *database.DBRepositories) *AppServices {
	as := &AppServices{}
	as.AuthorizationService = appservices.NewAuthorizationService(dbRepositories.AuthorizationDBRepository)
	return as
}
