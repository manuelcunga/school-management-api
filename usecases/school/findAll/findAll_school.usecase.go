package usecase

import (
	"github.com/manuelcunga/school-management-api/domain/repository"
	"github.com/manuelcunga/school-management-api/usecases/school/dtos"
)

type FindAllSchoolUsecase struct {
	schoolRepo repository.ISchoolRepository
}

func NewFindAllUSchoolUsecases(repo repository.ISchoolRepository) *FindAllSchoolUsecase {
	return &FindAllSchoolUsecase{schoolRepo: repo}
}

func (s FindAllSchoolUsecase) Execute() ([]*dtos.SchoolOutput, error) {
	users, err := s.schoolRepo.FindAll()

	if err != nil {
		return nil, err
	}

	var schoolOutputs []*dtos.SchoolOutput

	for _, school := range users {
		schoolOutputs = append(schoolOutputs, &dtos.SchoolOutput{
			ID:         school.ID,
			Name:       school.Name,
			Email:      school.Email,
			Province:   school.Province,
			RoomNumber: school.RoomNumber,
			CreatedAt:  school.CreatedAt,
			UpdatedAt:  school.UpdatedAt,
		})

	}

	return schoolOutputs, nil

}
