package paciente

import (
	"fmt"
	"strconv"

	"github.com/Hernandsv01/final-go.git/internal/domain"
)

type Service interface {
	Create(d domain.Paciente) (domain.Paciente, error)
	ReadAll() []domain.Paciente
	Read(dni int) (domain.Paciente, error)
	Update(dni string, d domain.Paciente, functionType string) error
	Delete(dni int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(d domain.Paciente) (domain.Paciente, error) {
	p, err := s.r.Create(d)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}

func (s *service) ReadAll() []domain.Paciente {
	return s.r.GetAll()
}

func (s *service) Read(dni int) (domain.Paciente, error) {
	return s.r.Get(dni)
}

func (s *service) Update(dni string, d domain.Paciente, functionType string) error {
	dniInt, err := strconv.Atoi(dni)
	if err != nil {
		return err
	}
	d.DNI = dniInt

	if functionType == "put" {
		return s.r.Put(d)
	} else if functionType == "patch" {
		return s.r.Patch(d)
	} else {
		return fmt.Errorf("Something went VERY wrong, contact developer")
	}
}

func (s *service) Delete(dni int) error {
	err := s.r.Delete(dni)
	return err
}
