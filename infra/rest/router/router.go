package router

import (
	"github.com/labstack/echo/v4"
	_ "github.com/manuelcunga/school-management-api/api"
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
	router.Echo.GET("/docs/*", echoSwagger.WrapHandler)
	return router.Echo

}

// @title School Management API
// @version 1.0
// @description This is a sample server sSchool Management API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /
func (router *RouterImpl) home() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(200, "School manegment version 1.0")
	}
}
