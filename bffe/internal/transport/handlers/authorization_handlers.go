package handlers

import (
	"github.com/gofiber/fiber/v3"
	appservices "github.com/o-rensa/jalikod-backend/bffe/internal/application/services"
	domainutilities "github.com/o-rensa/jalikod-backend/bffe/internal/domain/utility"
	handlererrors "github.com/o-rensa/jalikod-backend/bffe/internal/transport"
)

type AuthorizationHandlers struct {
	authorizationService *appservices.AuthorizationService
}

func NewAuthorizationHandlers(authorizationService *appservices.AuthorizationService) *AuthorizationHandlers {
	h := &AuthorizationHandlers{}
	h.authorizationService = authorizationService
	return h
}

func (ah *AuthorizationHandlers) RegisterRoutes(route fiber.Router) {
	authGroup := route.Group("/auth")
	authGroup.Post("/login", func(c fiber.Ctx) error { return ah.loginHandler(c) })
	authGroup.Get("/register", func(c fiber.Ctx) error { return ah.registerHandler(c) })
}

func (ah *AuthorizationHandlers) loginHandler(c fiber.Ctx) error {
	return c.SendString("Log in successful")
}

func (ah *AuthorizationHandlers) registerHandler(c fiber.Ctx) error {
	var rb RegisterRequest
	hn := "registerHandler"

	// check if request can bind to dto
	if err := c.Bind().Body(&rb); err != nil {
		handlererrors.BindRequestToDtoError(hn)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// validate request
	if err := domainutilities.Validate.Struct(rb); err != nil {
		handlererrors.ValidatorError(hn)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return nil
}
