package usecases

import (
	"fmt"

	"github.com/manuelcunga/school-management-api/domain/repository"
	"github.com/manuelcunga/school-management-api/usecases/school/dtos"
	util "github.com/manuelcunga/school-management-api/utils"
)

type UpdateSchoolUsecase struct {
	IschoolRepo repository.ISchoolRepository
}

func NewUpdateSchoolUsecase(repo repository.ISchoolRepository) *UpdateSchoolUsecase {
	return &UpdateSchoolUsecase{IschoolRepo: repo}
}

func (s UpdateSchoolUsecase) Execute(schoolID string, data *dtos.UpdateSchoolInput) (*dtos.SchoolOutput, error) {
	school, err := s.IschoolRepo.FindByID(schoolID)
	if err != nil {
		return nil, fmt.Errorf(string(util.NotFoundSchool))
	}

	if err = school.Update(data.Name, data.Email, data.Province, data.RoomNumber); err != nil {
		return nil, err
	}

	newSchool, err := s.IschoolRepo.Update(school)
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
