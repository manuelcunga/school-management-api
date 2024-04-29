package usecases

import (
	"fmt"

	"github.com/manuelcunga/school-management-api/domain/repository"
	util "github.com/manuelcunga/school-management-api/utils"
)

type DeleteSchoolUsecase struct {
	ISchoolRepository repository.ISchoolRepository
}

func NewDeleteSchoolUsecase(schoolRepo repository.ISchoolRepository) *DeleteSchoolUsecase {
	return &DeleteSchoolUsecase{ISchoolRepository: schoolRepo}
}

func (s DeleteSchoolUsecase) Delete(schoolID string) error {
	school, err := s.ISchoolRepository.FindByID(schoolID)
	fmt.Println(school)

	if err != nil {
		return fmt.Errorf(string(util.NotFoundSchool))
	}

	if err = s.ISchoolRepository.Delete(school.ID); err != nil {
		return fmt.Errorf(string(util.InternalServerError))
	}

	return nil

}
