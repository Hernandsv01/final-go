package paciente

import (
	"github.com/Hernandsv01/final-go.git/internal/domain"
	"github.com/Hernandsv01/final-go.git/pkg/store"
)

type Repository interface {
	Create(d domain.Paciente) (domain.Paciente, error)
	GetAll() []domain.Paciente
	Get(dni int) (domain.Paciente, error)
	Patch(d domain.Paciente) error
	Put(d domain.Paciente) error
	Delete(dni int) error
}

type repository struct {
	storage store.PacienteStoreInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.PacienteStoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) Create(d domain.Paciente) (domain.Paciente, error) {
	err := r.storage.Create(d)
	if err != nil {
		return domain.Paciente{}, err
	}

	return d, nil
}

func (r *repository) GetAll() []domain.Paciente {
	resList, err := r.storage.ReadAll()
	if err != nil {
		return make([]domain.Paciente, 0)
	}

	return resList
}

func (r *repository) Get(dni int) (domain.Paciente, error) {
	res, err := r.storage.Read(dni)
	if err != nil {
		return domain.Paciente{}, err
	}

	return res, nil
}

func (r *repository) Put(d domain.Paciente) error {
	return r.storage.UpdateFull(d)
}

func (r *repository) Patch(d domain.Paciente) error {
	return r.storage.Update(d)
}

func (r *repository) Delete(dni int) error {
	return r.storage.Delete(dni)
}