package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	database "github.com/manuelcunga/school-management-api/infra/db"
	"github.com/manuelcunga/school-management-api/infra/rest/dependency"
)

type Server struct {
	Router *echo.Echo
}

func NewServer(router *echo.Echo) *Server {
	return &Server{
		Router: router,
	}
}

func (server *Server) Start() {
	db := database.Connection()
	corsConfig := middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Permitir todas as origens (não recomendado para produção)
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}

	dependency.Dependency(db, server.Router)
	server.Router.Use(middleware.CORSWithConfig(corsConfig))
	// server.Router.Use(middleware.Logger())
	server.Router.Use(middleware.Recover())
	server.Router.Logger.Fatal(server.Router.Start(":8000"))
}
