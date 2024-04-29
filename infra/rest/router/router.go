package router

import (
	"github.com/labstack/echo/v4"
	"github.com/manuelcunga/school-management-api/infra/rest/router/interfaces"
	school_router "github.com/manuelcunga/school-management-api/infra/rest/school/router"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type RouterImpl struct {
	Echo *echo.Echo
	school_router.SchoolRouter
}

func NewRouter(router *RouterImpl) interfaces.IRouter {
	return &RouterImpl{
		Echo:         router.Echo,
		SchoolRouter: router.SchoolRouter,
	}
}

func (router *RouterImpl) Router() *echo.Echo {

	router.SchoolRouter.Router()
	router.Echo.GET("/", router.home())
	router.Echo.GET("/docs/*any", echoSwagger.WrapHandler)
	return router.Echo

}

// home is a simple handler to test our server
// @Summary Home
// @Description Home page of the API server of School manegment
// @Tags Home
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"Hello, Angola!"
// @Router / [get]

func (router *RouterImpl) home() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(200, "School manegment version 1.0")
	}
}
