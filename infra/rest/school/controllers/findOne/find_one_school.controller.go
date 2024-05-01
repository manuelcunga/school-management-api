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

// @Summary Find One  School
// @Description Find One School
// @Tags School
// @Accept  json
// @Produce  json
// @Param id path string true "School ID"
// @Success 200 {object}  dtos.SchoolOutput
// @Failure 500 {object}  messageserrors.MessageResponse
// @Router /api/v1/schools/{id} [get]
func (crl FindOneSchoolController) Handler(ctx echo.Context) error {
	id := ctx.Param("id")

	school, err := crl.findoneSchoolUsecase.Execute(id)

	if err != nil {
		return ctx.JSON(http.StatusNotFound, messageserrors.MessageResponse{Message: err.Error()})

	}

	return ctx.JSON(http.StatusOK, school)
}
