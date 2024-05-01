package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	messageserrors "github.com/manuelcunga/school-management-api/infra/rest/messages_errors"
	"github.com/manuelcunga/school-management-api/usecases/school/dtos"
	usecases "github.com/manuelcunga/school-management-api/usecases/school/update"
)

type UpdateSchoolController struct {
	uodateSchoolUsecase usecases.UpdateSchoolUsecase
}

func NewUpdateSchoolController(ScholUsecase usecases.UpdateSchoolUsecase) *UpdateSchoolController {
	return &UpdateSchoolController{uodateSchoolUsecase: ScholUsecase}
}

// @Summary Update School
// @Description Update School
// @Tags School
// @Accept  json
// @Produce  json
// @Param id path string true "School ID"
// @Param School body dtos.UpdateSchoolInput true "School"
// @Success 200 {object} dtos.SchoolOutput
// @Failure 500 {object} messageserrors.MessageResponse
// @Router /api/v1/schools/{id} [put]
func (s UpdateSchoolController) Handler(ctx echo.Context) error {
	id := ctx.Param("id")

	var school dtos.UpdateSchoolInput
	if err := ctx.Bind(&school); err != nil {
		return ctx.JSON(http.StatusBadRequest, messageserrors.MessageResponse{Message: err.Error()})
	}

	schoolUpdated, err := s.uodateSchoolUsecase.Execute(id, &school)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, messageserrors.MessageResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, schoolUpdated)

}
