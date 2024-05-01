package entities_test

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/manuelcunga/school-management-api/domain/entities"

	"github.com/stretchr/testify/assert"
)

func TestSchool(t *testing.T) {
	name := "John Doe"
	email := "school@example.com"
	province := "luanda"
	roomNumber := "912345678"
	school, err := entities.NewSchool(name, email, province, roomNumber)

	assert.Nil(t, err, "Expected nil, got error %v", err)
	assert.NotNil(t, school, "Expected user not to be nil")
	assert.Equal(t, name, school.Name, "Expected %s, got %s", name, school.Name)
	assert.Equal(t, email, school.Email, "Expected %s, got %s", email, school.Email)
	assert.Equal(t, province, school.Province, "Expected %s, got %s", province, school.Province)
	assert.Equal(t, roomNumber, school.RoomNumber, "Expected %s, got %s", roomNumber, school.RoomNumber)
	assert.NotEmpty(t, school.ID, "Expected user ID to be non-empty")
}

func TestUserUpdate(t *testing.T) {
	name := "John Doe"
	email := "school@example.com"
	province := "luanda"
	roomNumber := "912345678"
	school, err := entities.NewSchool(name, email, province, roomNumber)
	assert.Nil(t, err, "Expected nil, got error %v", err)

	newName := "teste"
	newProvince := "benguela"
	newRoomNumber := "12"
	err = school.Update(newName, faker.Email(), newProvince, newRoomNumber)

	assert.Nil(t, err, "Expected nil, got error %v", err)
	assert.Equal(t, newName, school.Name, "Expected %s, got %s", newName, school.Name)
}
