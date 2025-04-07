package api

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

type APIServer struct {
	port string
}

func NewAPIServer(port int) *APIServer {
	return &APIServer{
		port: fmt.Sprintf(":%d", port),
	}
}

func (s *APIServer) Run() error {
	mux := fiber.New()

	return mux.Listen(s.port)
}
