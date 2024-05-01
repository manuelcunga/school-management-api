package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	messageserrors "github.com/manuelcunga/school-management-api/infra/rest/messages_errors"
	usecases "github.com/manuelcunga/school-management-api/usecases/school/create"
	"github.com/manuelcunga/school-management-api/usecases/school/dtos"
)

type CreateSchoolController struct {
	createSchoolUsecae usecases.CreateSchoolUsecases
}

func NewCreateSchoolController(schoolUsecase usecases.CreateSchoolUsecases) *CreateSchoolController {
	return &CreateSchoolController{createSchoolUsecae: schoolUsecase}
}

// @Summary Create School
// @Description Create school
// @Tags School
// @Accept  json
// @Produce  json
// @Param school body dtos.CreateSchoolInput true "School"
// @Success 201 {object} dtos.SchoolOutput
// @Failure 500 {object} messageserrors.MessageResponse
// @Router /api/v1/schools [post]
func (ctrl *CreateSchoolController) Handler(ctx echo.Context) error {
	var schoolDTO dtos.CreateSchoolInput

	if err := ctx.Bind(&schoolDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, messageserrors.MessageResponse{Message: err.Error()})
	}

	output, err := ctrl.createSchoolUsecae.Execute(&schoolDTO)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, messageserrors.MessageResponse{Message: err.Error()})

	}

	return ctx.JSON(http.StatusCreated, output)
}
