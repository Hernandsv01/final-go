package dentista

import (
	"github.com/Hernandsv01/final-go.git/internal/domain"
	"github.com/Hernandsv01/final-go.git/pkg/store"
)

type Repository interface {
	Create(p domain.Dentista) (domain.Dentista, error)
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