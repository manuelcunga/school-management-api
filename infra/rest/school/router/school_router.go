package school_router

import (
	"os"

	"github.com/labstack/echo/v4"
	controllers "github.com/manuelcunga/school-management-api/infra/rest/school/controllers/create"
	deleteSchoolCtrl "github.com/manuelcunga/school-management-api/infra/rest/school/controllers/delete"
	findAllSchoolCtrl "github.com/manuelcunga/school-management-api/infra/rest/school/controllers/findAll"
	findOndSchoolCtrl "github.com/manuelcunga/school-management-api/infra/rest/school/controllers/findOne"
	updateSchoolCtrl "github.com/manuelcunga/school-management-api/infra/rest/school/controllers/update"
	uploadFileCtrl "github.com/manuelcunga/school-management-api/infra/rest/school/controllers/upload"
)

type SchoolRouter interface {
	Router() *echo.Echo
}

type SchoolRouterImpl struct {
	Echo              *echo.Echo
	CreateSchoolCtrl  controllers.CreateSchoolController
	FindOneSchoolCtrl findOndSchoolCtrl.FindOneSchoolController
	FindAllSchoolCtrl findAllSchoolCtrl.FindAllSchoolController
	DeleteSchoolCtrl  deleteSchoolCtrl.DeleteSchoolController
	UpdateSchoolCtrl  updateSchoolCtrl.UpdateSchoolController
	UploadSchoolCtrl  uploadFileCtrl.UploadController
}

func NewSchoolRouter(schoolRouter *SchoolRouterImpl) SchoolRouter {
	return &SchoolRouterImpl{
		Echo:              schoolRouter.Echo,
		CreateSchoolCtrl:  schoolRouter.CreateSchoolCtrl,
		FindOneSchoolCtrl: schoolRouter.FindOneSchoolCtrl,
		FindAllSchoolCtrl: schoolRouter.FindAllSchoolCtrl,
		DeleteSchoolCtrl:  schoolRouter.DeleteSchoolCtrl,
		UpdateSchoolCtrl:  schoolRouter.UpdateSchoolCtrl,
		UploadSchoolCtrl:  schoolRouter.UploadSchoolCtrl,
	}
}

func (router *SchoolRouterImpl) Router() *echo.Echo {
	v1 := router.Echo.Group(os.Getenv("BASE_PATH"))

	{
		school := v1.Group("/schools")
		{
			school.POST("", func(c echo.Context) error { return router.CreateSchoolCtrl.Handler(c) })
			school.GET("", func(c echo.Context) error { return router.FindAllSchoolCtrl.Handler(c) })
			school.GET("/:id", func(c echo.Context) error { return router.FindOneSchoolCtrl.Handler(c) })
			school.PUT("/:id", func(c echo.Context) error { return router.UpdateSchoolCtrl.Handler(c) })
			school.DELETE("/:id", func(c echo.Context) error { return router.DeleteSchoolCtrl.Handler(c) })
			school.POST("/upload", func(c echo.Context) error { return router.UploadSchoolCtrl.Handler(c) })

		}
	}

	return router.Echo

}
