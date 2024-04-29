package repository

import (
	"github.com/manuelcunga/school-management-api/domain/entities"
	"github.com/manuelcunga/school-management-api/domain/repository"
	"gorm.io/gorm"
)

type SchoolRepository struct {
	Db *gorm.DB
}

func NewSchoolRepository(db *gorm.DB) repository.ISchoolRepository {
	return &SchoolRepository{Db: db}
}

func (s *SchoolRepository) Create(entity *entities.School) error {
	return s.Db.Create(entity).Error
}

func (s *SchoolRepository) FindByEmail(email string) (bool, error) {
	var count int64
	err := s.Db.Model(&entities.School{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err

}

func (s *SchoolRepository) Update(entity *entities.School) (*entities.School, error) {
	err := s.Db.Save(entity).Error
	return entity, err
}

func (s *SchoolRepository) Delete(id string) error {
	return s.Db.Delete(&entities.School{}, "id = ?", id).Error
}

func (s *SchoolRepository) FindAll() ([]*entities.School, error) {
	var school []*entities.School
	err := s.Db.Find(&school).Error
	return school, err
}

func (s *SchoolRepository) FindByID(id string) (*entities.School, error) {
	school := entities.School{}
	err := s.Db.First(&school, "id = ?", id).Error
	return &school, err

}
