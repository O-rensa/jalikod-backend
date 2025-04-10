package api

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3"
	transport "github.com/o-rensa/jalikod-backend/bffe/internal/transport/handlers/Authorization"
)

type APIServer struct {
	port string
	db   *sql.DB
}

func NewAPIServer(port string, db *sql.DB) *APIServer {
	return &APIServer{
		port: fmt.Sprintf(":%s", port),
		db:   db,
	}
}

func (s *APIServer) Run() error {
	muxConfig := fiber.Config{
		AppName: os.Getenv("APP_NAME"),
	}
	mux := fiber.New(muxConfig)
	groupedRoutes := mux.Group("/v1")

	authorizationHandler := transport.NewAuthorizationHandlers()
	authorizationHandler.RegisterRoutes(groupedRoutes)

	return mux.Listen(s.port)
}
