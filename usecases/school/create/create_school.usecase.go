package usecases

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/manuelcunga/school-management-api/domain/entities"
	"github.com/manuelcunga/school-management-api/domain/repository"
	"github.com/manuelcunga/school-management-api/usecases/school/dtos"
	util "github.com/manuelcunga/school-management-api/utils"
)

type CreateSchoolUsecases struct {
	createSchoolRepo repository.ISchoolRepository
	provinces        map[string][]string
}

type Provincia struct {
	Nome    string `json:"nome"`
	Capital string `json:"capital"`
}

type Angola struct {
	Provincias []Provincia `json:"Provincias"`
}

func NewCreateUserUseCase(schoolRepo repository.ISchoolRepository) *CreateSchoolUsecases {
	provinces, err := loadProvinces()
	if err != nil {
		panic(fmt.Errorf("failed to load provinces: %v", err))
	}

	return &CreateSchoolUsecases{
		createSchoolRepo: schoolRepo,
		provinces:        provinces,
	}
}

func loadProvinces() (map[string][]string, error) {
	data, err := os.ReadFile("config/provincias.json")
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}

	var angolaData map[string]Angola

	if err := json.Unmarshal(data, &angolaData); err != nil {
		log.Fatalf("Erro ao decodificar JSON: %v", err)
	}

	angola, ok := angolaData["Angola"]
	if !ok {
		log.Fatal("Não foi possível encontrar os dados de Angola")
	}

	provinces := make(map[string][]string)
	for _, provincia := range angola.Provincias {
		provinces["Angola"] = append(provinces["Angola"], provincia.Nome)
	}

	return provinces, nil
}

func (school *CreateSchoolUsecases) Execute(input *dtos.CreateSchoolInput) (*dtos.SchoolOutput, error) {
	if !school.provinceExists(input.Province) {
		return nil, fmt.Errorf("province does not exist")
	}

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

func (school *CreateSchoolUsecases) provinceExists(province string) bool {
	provincesOfAngola, ok := school.provinces["Angola"]
	if !ok {
		return false
	}

	for _, p := range provincesOfAngola {
		if p == province {
			return true
		}
	}

	return false
}
