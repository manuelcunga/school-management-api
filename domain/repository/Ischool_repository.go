package repository

import "github.com/manuelcunga/school-management-api/domain/entities"

type ISchoolRepository interface {
	Create(entity *entities.School) error
	Update(entity *entities.School) (*entities.School, error)
	Delete(id string) error
	FindByID(id string) (*entities.School, error)
	FindAll() ([]*entities.School, error)
	FindByEmail(email string) (bool, error)
}
