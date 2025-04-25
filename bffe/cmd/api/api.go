package api

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/o-rensa/jalikod-backend/bffe/internal/application"
	"github.com/o-rensa/jalikod-backend/bffe/internal/infrastructure/database"
	"github.com/o-rensa/jalikod-backend/bffe/internal/transport/handlers"
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
	// database repository implementation
	dbr := database.NewDBRepositories(s.db)
	// app services
	as := application.NewApplicationServices(dbr)
	muxConfig := fiber.Config{
		AppName: os.Getenv("APP_NAME"),
	}
	mux := fiber.New(muxConfig)
	groupedRoutes := mux.Group("/api/v1")

	authenticationHandler := handlers.NewAuthenticationHandlers(as.AuthenticationService)
	authenticationHandler.RegisterRoutes(groupedRoutes)

	return mux.Listen(s.port)
}
