package turno

import (
	"github.com/Hernandsv01/final-go.git/internal/domain"
	"github.com/Hernandsv01/final-go.git/pkg/store"
)

type Repository interface {
	Create(d domain.Turno) error
	GetAll() []domain.Turno
	Get(id int) (domain.Turno, error)
	Patch(d domain.Turno) error
	Put(d domain.Turno) error
	Delete(id int) error
	GetByDni(dni int) ([]domain.Turno, error)
}

type repository struct {
	storage store.TurnoStoreInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.TurnoStoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) Create(t domain.Turno) error {
	return r.storage.Create(t)
}

func (r *repository) GetAll() []domain.Turno {
	turnoList, _ := r.storage.ReadAll()
	return turnoList
}

func (r *repository) Get(id int) (domain.Turno, error) {
	return r.storage.Read(id)
}

func (r *repository) Put(t domain.Turno) error {
	return r.storage.UpdateFull(t)
}

func (r *repository) Patch(t domain.Turno) error {
	return r.storage.Update(t)
}

func (r *repository) Delete(id int) error {
	return r.storage.Delete(id)
}

func (r *repository) GetByDni(dni int) ([]domain.Turno, error) {
	return r.storage.GetByDni(dni)
}