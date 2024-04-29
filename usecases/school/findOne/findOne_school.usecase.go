package usecases

import (
	"github.com/manuelcunga/school-management-api/domain/repository"
	"github.com/manuelcunga/school-management-api/usecases/school/dtos"
)

type FindOneSchoolUsecase struct {
	IschollRepository repository.ISchoolRepository
}

func NewFindOneSchoolUsecase(schoolRepo repository.ISchoolRepository) *FindOneSchoolUsecase {
	return &FindOneSchoolUsecase{IschollRepository: schoolRepo}
}

func (s FindOneSchoolUsecase) Execute(userId string) (*dtos.SchoolOutput, error) {
	school, err := s.IschollRepository.FindByID(userId)

	if err != nil {
		return nil, err
	}

	return &dtos.SchoolOutput{
		ID:         school.ID,
		Name:       school.Name,
		Email:      school.Email,
		Province:   school.Province,
		RoomNumber: school.RoomNumber,
		CreatedAt:  school.CreatedAt,
		UpdatedAt:  school.UpdatedAt,
	}, nil
}
