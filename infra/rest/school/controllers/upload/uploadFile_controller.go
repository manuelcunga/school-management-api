package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
	messageserrors "github.com/manuelcunga/school-management-api/infra/rest/messages_errors"
	usecases "github.com/manuelcunga/school-management-api/usecases/school/upload"
)

type UploadController struct {
	UploadUseCase usecases.UploadUseCase
}

func NewUploadController(uploadUseCase usecases.UploadUseCase) *UploadController {
	return &UploadController{
		UploadUseCase: uploadUseCase,
	}
}

// @Summary Upload Excel file School
// @Description Register School from Import excel file
// @Tags School
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Excel file to upload (.xlsx)"
// @Success 200 {object} dtos.SchoolOutput
// @Failure 400 {object} messageserrors.MessageResponse
// @Router /api/v1/schools/upload [post]
func (ctrl *UploadController) Handler(ctx echo.Context) error {
	file, err := ctx.FormFile("file")
	if err != nil {

		return ctx.JSON(http.StatusBadRequest, messageserrors.MessageResponse{Message: "Failed to get file from request"})

	}

	if filepath.Ext(file.Filename) != ".xlsx" {

		return ctx.JSON(http.StatusBadRequest, messageserrors.MessageResponse{Message: "Only Excel files are allowed"})

	}

	done := make(chan struct{})

	go func() {
		defer close(done)
		err := ctrl.UploadUseCase.Execute(file)
		if err != nil {
			fmt.Println("Failed to upload file:", err)
		}

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, messageserrors.MessageResponse{Message: err.Error()})
		}

	}()

	<-done

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "File upload successfull",
	})
}
