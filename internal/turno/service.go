package turno

import (
	"fmt"
	"strconv"

	"github.com/Hernandsv01/final-go.git/internal/dentista"
	"github.com/Hernandsv01/final-go.git/internal/domain"
	"github.com/Hernandsv01/final-go.git/internal/paciente"
)

type Service interface {
	Create(d domain.Turno) error
	ReadAll() []domain.Turno
	Read(id int) (domain.Turno, error)
	Update(id string, d domain.Turno, functionType string) error
	Delete(id int) error
	GetByDni(dni int) ([]domain.Turno, error)
}

type service struct {
	r Repository
	sd dentista.Service
	sp paciente.Service
}

func NewService(r Repository, sd dentista.Service, sp paciente.Service) Service {
	return &service{r, sd ,sp}
}

func (s *service) Create(t domain.Turno) error {
	/*
		Verificar que estén los campos obligatorios
		Verificar que el dni exista
		Verificar que la matricula exista
		Crear
	*/
	if t.Dentista.Matricula == 0 || t.Paciente.DNI == 0 {
		return fmt.Errorf("Se necesita que estén tanto la matricula del dentista como el dni del paciente para crear un turno")
	}
	todoCorrectoDentista := s.sd.Exists(t.Dentista.Matricula)
	todoCorrectoPaciente := s.sp.Exists(t.Paciente.DNI)
	if !todoCorrectoDentista || !todoCorrectoPaciente {
		return fmt.Errorf("No existe ningún paciente o ningún dentista con esos datos (Service): %v %v", todoCorrectoDentista, todoCorrectoPaciente)
	}

	err := s.r.Create(t)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ReadAll() []domain.Turno {
	return s.r.GetAll()
}

func (s *service) Read(id int) (domain.Turno, error) {
	return s.r.Get(id)
}

func (s *service) Update(id string, t domain.Turno, functionType string) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	t.Id = idInt

	if functionType == "put" {
		return s.r.Put(t)
	} else if functionType == "patch" {
		return s.r.Patch(t)
	} else {
		return fmt.Errorf("Something went VERY wrong, contact developer")
	}
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	return err
}

func (s *service) GetByDni(dni int) ([]domain.Turno, error) {
	return s.r.GetByDni(dni)

}