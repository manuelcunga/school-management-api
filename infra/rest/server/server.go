package server

import (
	"github.com/labstack/echo/v4"
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
	dependency.Dependency(db, server.Router)
	// Middleware
	// server.Router.Use(middleware.Logger())
	// server.Router.Use(middleware.Recover())
	server.Router.Logger.Fatal(server.Router.Start(":8000"))
}
