package usecases

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/manuelcunga/school-management-api/domain/entities"
	"github.com/manuelcunga/school-management-api/domain/repository"
	"github.com/manuelcunga/school-management-api/usecases/school/dtos"
	util "github.com/manuelcunga/school-management-api/utils"
	"github.com/manuelcunga/school-management-api/utils/firebase"
	"github.com/tealeg/xlsx"
)

type UploadUseCase struct {
	schoolRepo repository.ISchoolRepository
}

func NewUploadUseCase(schoolRepo repository.ISchoolRepository) *UploadUseCase {
	return &UploadUseCase{
		schoolRepo: schoolRepo,
	}
}

func (uc *UploadUseCase) Execute(file *multipart.FileHeader) error {
	// Abrir o arquivo enviado pelo cliente
	fileReader, err := file.Open()
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer fileReader.Close()

	tempFile, err := os.CreateTemp("", "files-*.xlsx")
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	if _, err := io.Copy(tempFile, fileReader); err != nil {
		return fmt.Errorf("failed to copy file content to temporary file: %v", err)
	}
	tempFile.Close()

	xlFile, err := xlsx.OpenFile(tempFile.Name())
	if err != nil {
		return fmt.Errorf("failed to open Excel file: %w", err)
	}

	// Ignorar a primeira linha (cabeçalho)
	firstRowSkipped := false

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			// Ignorar a primeira linha
			if !firstRowSkipped {
				firstRowSkipped = true
				continue
			}

			cells := row.Cells
			if len(cells) < 4 {
				continue
			}

			schoolFromExcel := &dtos.CreateSchoolInput{
				Name:       cells[0].Value,
				Email:      cells[1].Value,
				Province:   cells[3].Value,
				RoomNumber: cells[2].Value,
			}

			newSchool, err := entities.NewSchool(schoolFromExcel.Name, schoolFromExcel.Email, schoolFromExcel.Province, schoolFromExcel.RoomNumber)
			if err != nil {
				return fmt.Errorf("failed to create school: %v", err)
			}

			exists, err := uc.schoolRepo.FindByEmail(newSchool.Email)
			if err != nil {
				return fmt.Errorf("failed to check school existence: %v", err)
			}
			if exists {
				return fmt.Errorf(string(util.SchoolAlredayExists))

			}

			if err := uc.schoolRepo.Create(newSchool); err != nil {
				return fmt.Errorf("failed to create school: %v", err)
			}
		}
	}

	client := firebase.GetStorageClient()

	bucket, err := client.Bucket(os.Getenv("FIREBASE_BUCKET_NAME"))
	if err != nil {
		return fmt.Errorf("error getting bucket: %v", err)
	}

	filename := file.Filename

	obj := bucket.Object(filename)

	wc := obj.NewWriter(context.Background())
	wc.ContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"

	if _, err := io.Copy(wc, fileReader); err != nil {
		return fmt.Errorf("error copying file to Firebase Storage: %v", err)
	}

	if err := wc.Close(); err != nil {
		return fmt.Errorf("error closing writer: %v", err)
	}

	return nil
}
