package application

import (
	appservices "github.com/o-rensa/jalikod-backend/bffe/internal/application/services"
	"github.com/o-rensa/jalikod-backend/bffe/internal/infrastructure/database"
)

type AppServices struct {
	AuthenticationService *appservices.AuthenticationService
}

func NewApplicationServices(dbRepositories *database.DBRepositories) *AppServices {
	as := &AppServices{}
	as.AuthenticationService = appservices.NewAuthenticationService(dbRepositories.AuthenticationDBRepository)
	return as
}
