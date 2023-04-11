package dentista

import (
	"github.com/Hernandsv01/final-go.git/internal/domain"
	"github.com/Hernandsv01/final-go.git/pkg/store"
)

type Repository interface {
	Create(d domain.Dentista) (domain.Dentista, error)
	GetAll() []domain.Dentista
	Get(matricula int) (domain.Dentista, error)
	Patch(d domain.Dentista) error
	Put(d domain.Dentista) error
	Delete(matricula int) error
}

type repository struct {
	storage store.DentistaStoreInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.DentistaStoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) Create(d domain.Dentista) (domain.Dentista, error) {
	err := r.storage.Create(d)
	if err != nil {
		return domain.Dentista{}, err
	}

	return d, nil
}

func (r *repository) GetAll() []domain.Dentista {
	resList, err := r.storage.ReadAll()
	if err != nil {
		return make([]domain.Dentista, 0)
	}

	return resList
}

func (r *repository) Get(matricula int) (domain.Dentista, error) {
	res, err := r.storage.Read(matricula)
	if err != nil {
		return domain.Dentista{}, err
	}

	return res, nil
}

func (r *repository) Put(d domain.Dentista) error {
	return r.storage.UpdateFull(d)
}

func (r *repository) Patch(d domain.Dentista) error {
	return r.storage.Update(d)
}

func (r *repository) Delete(matricula int) error {
	return r.storage.Delete(matricula)
}