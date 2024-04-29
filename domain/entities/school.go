package entities

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type School struct {
	ID         string    `json:"id" gorm:"type:uuid;primaryKey" valid:"-"`
	Name       string    `json:"name" valid:"required~ O nome  é obrigatório."`
	Email      string    `json:"email" valid:"required~ O E-mail é obrigatório."`
	RoomNumber string    `json:"room_number" valid:"required~ O número de sala é obrigatório."`
	Province   string    `json:"province" valid:"required~ A província é obrigatório."`
	CreatedAt  time.Time `json:"created_at" valid:"-" gorm:"type:timestamp without time zone"`
	UpdatedAt  time.Time `json:"updated_at" valid:"-"`
	DeletedAt  time.Time `json:"deleted_at"  valid:"-"`
}

func NewSchool(name, email, province, roomNumber string) (*School, error) {
	school := &School{
		Name:       name,
		Email:      email,
		RoomNumber: roomNumber,
		Province:   province,
	}

	school.ID = uuid.NewV4().String()
	school.CreatedAt = time.Now()

	if err := school.isValid(); err != nil {
		return nil, err
	}

	return school, nil
}

func (school *School) isValid() error {
	_, err := govalidator.ValidateStruct(school)
	if err != nil {
		return err
	}
	return nil
}

func (school *School) SetUpdatedAt() {
	school.UpdatedAt = time.Now()
}

func (school *School) Update(name, email, province, roomNumber string) error {

	if name != "" {
		school.Name = name
	}

	if email != "" {
		school.Email = email
	}

	if province != "" {
		school.Province = province
	}

	if roomNumber != "" {
		school.RoomNumber = roomNumber
	}

	school.SetUpdatedAt()

	return nil
}
