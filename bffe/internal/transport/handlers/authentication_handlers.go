package handlers

import (
	"github.com/gofiber/fiber/v3"
	appservices "github.com/o-rensa/jalikod-backend/bffe/internal/application/services"
	domainutilities "github.com/o-rensa/jalikod-backend/bffe/internal/domain/utility"
	handlererrors "github.com/o-rensa/jalikod-backend/bffe/internal/transport"
)

type AuthenticationHandlers struct {
	authenticationService *appservices.AuthenticationService
}

func NewAuthenticationHandlers(authenticationService *appservices.AuthenticationService) *AuthenticationHandlers {
	h := &AuthenticationHandlers{}
	h.authenticationService = authenticationService
	return h
}

func (ah *AuthenticationHandlers) RegisterRoutes(route fiber.Router) {
	authGroup := route.Group("/auth")
	authGroup.Post("/login", func(c fiber.Ctx) error { return ah.loginHandler(c) })
	authGroup.Post("/register", func(c fiber.Ctx) error { return ah.registerHandler(c) })
}

func (ah *AuthenticationHandlers) loginHandler(c fiber.Ctx) error {
	return c.SendString("Log in successful")
}

func (ah *AuthenticationHandlers) registerHandler(c fiber.Ctx) error {
	var req appservices.RegisterRequest
	hn := "registerHandler"

	// check if request can bind to dto
	if err := c.Bind().Body(&req); err != nil {
		handlererrors.BindRequestToDtoError(hn)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// validate request
	if err := domainutilities.Validate.Struct(req); err != nil {
		handlererrors.ValidatorError(hn)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var ctx = c.Context()

	// check if username already exists
	_, err := ah.authenticationService.CheckUsername(ctx, req.Username)
	if err == nil { // if no error, it means it found an existing username
		response := appservices.RegisterResponse{}
		response.Status = appservices.RegisterFail
		response.FailMessage = []appservices.RegisterFailure{1}
		c.JSON(response)
		return c.SendStatus(fiber.StatusOK)
	}

	// send to service
	stat, reg_stat, err := ah.authenticationService.RegisterUser(ctx, req)
	if err != nil {
		c.SendString(err.Error())
		c.SendStatus(stat)
	}
	c.JSON(reg_stat)
	return c.SendStatus(stat)
}
