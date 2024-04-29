package dependency

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Dependency(db *gorm.DB, echo *echo.Echo) *echo.Echo {

	DependencyRouter(db, echo)
	SchoolDependecy(db, echo)
	return echo
}
