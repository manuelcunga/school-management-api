package dependency

import (
	"github.com/labstack/echo/v4"
	"github.com/manuelcunga/school-management-api/infra/repository"
	"github.com/manuelcunga/school-management-api/infra/rest/router/interfaces"
	controllers "github.com/manuelcunga/school-management-api/infra/rest/school/controllers/create"
	deleteSchoolCtrl "github.com/manuelcunga/school-management-api/infra/rest/school/controllers/delete"
	findAllSchoolCtrl "github.com/manuelcunga/school-management-api/infra/rest/school/controllers/findAll"
	findOneSchoolCtrl "github.com/manuelcunga/school-management-api/infra/rest/school/controllers/findOne"
	updateSchoolCtrl "github.com/manuelcunga/school-management-api/infra/rest/school/controllers/update"
	uploadFileCtrl "github.com/manuelcunga/school-management-api/infra/rest/school/controllers/upload"
	school_router "github.com/manuelcunga/school-management-api/infra/rest/school/router"
	usecase "github.com/manuelcunga/school-management-api/usecases/school/create"
	deleteSchoolUsecase "github.com/manuelcunga/school-management-api/usecases/school/delete"
	findAllUsecases "github.com/manuelcunga/school-management-api/usecases/school/findAll"
	findOneSchoolUsecase "github.com/manuelcunga/school-management-api/usecases/school/findOne"
	updateSchoolUsecase "github.com/manuelcunga/school-management-api/usecases/school/update"
	uploadFileUsecase "github.com/manuelcunga/school-management-api/usecases/school/upload"

	"gorm.io/gorm"
)

func SchoolDependecy(db *gorm.DB, echo *echo.Echo) interfaces.IRouter {
	schoolRepo := repository.NewSchoolRepository(db)

	createSchoolUSecase := usecase.NewCreateUserUseCase(schoolRepo)
	findOneSchoolusecase := findOneSchoolUsecase.NewFindOneSchoolUsecase(schoolRepo)
	updateSchool := updateSchoolUsecase.NewUpdateSchoolUsecase(schoolRepo)
	deleteSchoolUsecase := deleteSchoolUsecase.NewDeleteSchoolUsecase(schoolRepo)
	uploadUseCase := uploadFileUsecase.NewUploadUseCase(schoolRepo)

	createSchoolController := controllers.NewCreateSchoolController(*createSchoolUSecase)
	findAllSchoolUsecase := findAllUsecases.NewFindAllUSchoolUsecases(schoolRepo)
	findAllUserController := findAllSchoolCtrl.NewFindAllSchoolController(*findAllSchoolUsecase)
	findOneSchoolController := findOneSchoolCtrl.NewFindOneSchoolController(*findOneSchoolusecase)
	updateSchoolController := updateSchoolCtrl.NewUpdateSchoolController(*updateSchool)
	deleteSchoolController := deleteSchoolCtrl.NewUpdateSchoolController(*deleteSchoolUsecase)
	uploadScholController := uploadFileCtrl.NewUploadController(*uploadUseCase)

	schoolRouter := &school_router.SchoolRouterImpl{
		Echo:              echo,
		CreateSchoolCtrl:  *createSchoolController,
		FindAllSchoolCtrl: *findAllUserController,
		FindOneSchoolCtrl: *findOneSchoolController,
		UpdateSchoolCtrl:  *updateSchoolController,
		DeleteSchoolCtrl:  *deleteSchoolController,
		UploadSchoolCtrl:  *uploadScholController,
	}

	return school_router.NewSchoolRouter(schoolRouter)
}
