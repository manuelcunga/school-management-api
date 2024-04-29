package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	messageserrors "github.com/manuelcunga/school-management-api/infra/rest/messages_errors"
	usecases "github.com/manuelcunga/school-management-api/usecases/school/findOne"
)

type FindOneSchoolController struct {
	findoneSchoolUsecase usecases.FindOneSchoolUsecase
}

func NewFindOneSchoolController(schoolUsecase usecases.FindOneSchoolUsecase) *FindOneSchoolController {
	return &FindOneSchoolController{findoneSchoolUsecase: schoolUsecase}
}

func (crl FindOneSchoolController) Handler(ctx echo.Context) error {
	id := ctx.Param("id")

	school, err := crl.findoneSchoolUsecase.Execute(id)

	if err != nil {
		return ctx.JSON(http.StatusNotFound, messageserrors.MessageResponse{Message: err.Error()})

	}

	return ctx.JSON(http.StatusOK, school)
}
