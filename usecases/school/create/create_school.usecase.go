package usecases

import (
	"fmt"

	"github.com/manuelcunga/school-management-api/domain/entities"
	"github.com/manuelcunga/school-management-api/domain/repository"
	"github.com/manuelcunga/school-management-api/usecases/school/dtos"
	util "github.com/manuelcunga/school-management-api/utils"
)

type CreateSchoolUsecases struct {
	createSchoolRepo repository.ISchoolRepository
}

func NewCreateUserUseCase(schoolRepo repository.ISchoolRepository) *CreateSchoolUsecases {
	return &CreateSchoolUsecases{
		createSchoolRepo: schoolRepo,
	}
}

func (school *CreateSchoolUsecases) Execute(input *dtos.CreateSchoolInput) (*dtos.SchoolOutput, error) {

	exists, err := school.createSchoolRepo.FindByEmail(input.Email)

	if err != nil {
		return nil, err
	}

	if exists {
		return nil, fmt.Errorf(string(util.SchoolAlredayExists))
	}

	newSchool, err := entities.NewSchool(input.Name, input.Email, input.Province, input.RoomNumber)
	if err != nil {
		return nil, err
	}

	err = school.createSchoolRepo.Create(newSchool)
	if err != nil {
		return nil, err
	}

	return &dtos.SchoolOutput{
		ID:         newSchool.ID,
		Name:       newSchool.Name,
		Email:      newSchool.Email,
		Province:   newSchool.Province,
		RoomNumber: newSchool.RoomNumber,
		CreatedAt:  newSchool.CreatedAt,
		UpdatedAt:  newSchool.UpdatedAt,
	}, nil
}
