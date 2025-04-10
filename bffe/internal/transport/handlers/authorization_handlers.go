package handlers

import (
	"github.com/gofiber/fiber/v3"
	appservices "github.com/o-rensa/jalikod-backend/bffe/internal/application/services"
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
	//var registerReq RegisterRequest
	message := ah.authorizationService.Hello()

	// if err := c.Bind().Body(&registerReq); err != nil {
	// return c.SendStatus(fiber.StatusBadRequest)
	// }
	return c.SendString(message)
}
