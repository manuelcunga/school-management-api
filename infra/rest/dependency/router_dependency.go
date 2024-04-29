package dependency

import (
	"github.com/labstack/echo/v4"
	"github.com/manuelcunga/school-management-api/infra/rest/router"
	"gorm.io/gorm"
)

func DependencyRouter(db *gorm.DB, echo *echo.Echo) *echo.Echo {

	app := router.RouterImpl{
		Echo:         echo,
		SchoolRouter: SchoolDependecy(db, echo),
	}

	router.NewRouter(&app).Router()
	return echo
}
