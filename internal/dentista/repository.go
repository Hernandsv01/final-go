package dentista

import (
	"github.com/Hernandsv01/final-go.git/internal/domain"
	"github.com/Hernandsv01/final-go.git/pkg/store"
)

type Repository interface {
	Create(p domain.Dentista) (domain.Dentista, error)
	GetAll() []domain.Dentista
	Get(matricula int) (domain.Dentista, error)
}

type repository struct {
	storage store.DentistaStoreInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.DentistaStoreInterface) Repository {
	return &repository{storage}
}

// Create agrega un nuevo producto
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