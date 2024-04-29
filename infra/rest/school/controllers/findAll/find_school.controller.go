package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	messageserrors "github.com/manuelcunga/school-management-api/infra/rest/messages_errors"
	usecase "github.com/manuelcunga/school-management-api/usecases/school/findAll"
)

type FindAllSchoolController struct {
	FindAllSchool usecase.FindAllSchoolUsecase
}

func NewFindAllSchoolController(schoolRepo usecase.FindAllSchoolUsecase) *FindAllSchoolController {
	return &FindAllSchoolController{FindAllSchool: schoolRepo}
}

func (s FindAllSchoolController) Handler(ctx echo.Context) error {
	schools, err := s.FindAllSchool.Execute()

	if err != nil {
		return ctx.JSON(http.StatusNotFound, messageserrors.MessageResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, schools)
}
