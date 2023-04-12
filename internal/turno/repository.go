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
	err := r.storage.Create(t)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAll() []domain.Turno {
	resList, err := r.storage.ReadAll()
	if err != nil {
		return make([]domain.Turno, 0)
	}

	return resList
}

func (r *repository) Get(id int) (domain.Turno, error) {
	res, err := r.storage.Read(id)
	if err != nil {
		return domain.Turno{}, err
	}

	return res, nil
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