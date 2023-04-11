package dentista

import (
	"github.com/Hernandsv01/final-go.git/internal/domain"
)

type Service interface {
	Create(d domain.Dentista) (domain.Dentista, error)
	ReadAll() []domain.Dentista
	Read(matricula int) (domain.Dentista, error)
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
/*
POST: agregar dentista.
GET: traer dentista por ID.
PUT: actualizar dentista.
PATCH: actualizar un dentista por alguno de sus campos.
DELETE: eliminar dentista.
*/