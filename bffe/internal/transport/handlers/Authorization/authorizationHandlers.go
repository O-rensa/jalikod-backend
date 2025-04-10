package transport

import (
	"github.com/gofiber/fiber/v3"
)

type AuthorizationHandlers struct {
}

func NewAuthorizationHandlers() *AuthorizationHandlers {
	handler := &AuthorizationHandlers{}
	return handler
}

func (ah *AuthorizationHandlers) RegisterRoutes(route fiber.Router) {
	authGroup := route.Group("/auth")
	authGroup.Post("/login", func(c fiber.Ctx) error { return ah.loginHandler(c) })
	authGroup.Post("/register", func(c fiber.Ctx) error { return ah.registerHandler(c) })
}

func (ah *AuthorizationHandlers) loginHandler(c fiber.Ctx) error {
	return c.SendString("Log in successful")
}

func (ah *AuthorizationHandlers) registerHandler(c fiber.Ctx) error {
	return c.SendString("Register successful")
}
