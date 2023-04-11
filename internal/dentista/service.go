package dentista

import (
	"strconv"

	"github.com/Hernandsv01/final-go.git/internal/domain"
)

type Service interface {
	Create(d domain.Dentista) (domain.Dentista, error)
	ReadAll() []domain.Dentista
	Read(matricula int) (domain.Dentista, error)
	Put(matricula string, d domain.Dentista) error
	Patch(matricula string, d domain.Dentista) error
	Delete(matricula int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(d domain.Dentista) (domain.Dentista, error) {
	p, err := s.r.Create(d)
	if err != nil {
		return domain.Dentista{}, err
	}
	return p, nil
}

func (s *service) ReadAll() []domain.Dentista {
	return s.r.GetAll()
}

func (s *service) Read(matricula int) (domain.Dentista, error) {
	return s.r.Get(matricula)
}

func (s *service) Put(matricula string, d domain.Dentista) error {
	matInt, err := strconv.Atoi(matricula)
	if err != nil { 
		return err
	}
	d.Matricula = matInt
	err = s.r.Put(d)
	return err
}

func (s *service) Patch(matricula string, d domain.Dentista) error {
	matInt, err := strconv.Atoi(matricula)
	if err != nil { 
		return err
	}
	d.Matricula = matInt
	err = s.r.Patch(d)
	return err
}

func (s *service) Delete(matricula int) error {
	err := s.r.Delete(matricula)
	return err
}

/*
POST: agregar dentista.
GET: traer dentista por ID.
PUT: actualizar dentista.
PATCH: actualizar un dentista por alguno de sus campos.
DELETE: eliminar dentista.
*/