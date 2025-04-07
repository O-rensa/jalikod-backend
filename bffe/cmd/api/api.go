package api

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v3"
)

type APIServer struct {
	port string
	db   *sql.DB
}

func NewAPIServer(port int, db *sql.DB) *APIServer {
	return &APIServer{
		port: fmt.Sprintf(":%d", port),
		db:   db,
	}
}

func (s *APIServer) Run() error {
	mux := fiber.New()

	return mux.Listen(s.port)
}
