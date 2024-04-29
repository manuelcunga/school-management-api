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
	}()

	<-done

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "File upload successfull",
	})
}
