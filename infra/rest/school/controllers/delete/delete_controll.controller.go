package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	messageserrors "github.com/manuelcunga/school-management-api/infra/rest/messages_errors"
	usecases "github.com/manuelcunga/school-management-api/usecases/school/delete"
)

type DeleteSchoolController struct {
	deleteSchoolController usecases.DeleteSchoolUsecase
}

func NewUpdateSchoolController(SchoolUsecase usecases.DeleteSchoolUsecase) *DeleteSchoolController {
	return &DeleteSchoolController{deleteSchoolController: SchoolUsecase}
}

// @Summary Delete School
// @Description Delete School
// @Tags School
// @Accept  json
// @Produce  json
// @Param id path string true "School ID"
// @Success 200 {object}  dtos.SchoolOutput
// @Failure 500 {object}  messageserrors.MessageResponse
// @Router /api/v1/schools/{id} [delete]
func (s DeleteSchoolController) Handler(ctx echo.Context) error {
	schoolID := ctx.Param("id")

	err := s.deleteSchoolController.Delete(schoolID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, messageserrors.MessageResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, messageserrors.MessageResponse{Message: "School deleted successfully"})
}
