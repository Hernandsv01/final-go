package store

import (
	"database/sql"

	"github.com/Hernandsv01/final-go.git/internal/domain"
)

type dentistaSqlStore struct {
	db *sql.DB
}

func NewDentistaSqlStore(db *sql.DB) DentistaStoreInterface {
	return &dentistaSqlStore{
		db: db,
	}
}


// Read devuelve un Dentista por su id
func (s *dentistaSqlStore) Read(id int) (domain.Dentista, error) {
	return domain.Dentista{}, nil
}

// Create agrega un nuevo Dentista
func (s *dentistaSqlStore) Create(product domain.Dentista) error {
 	return nil
}

// Update actualiza un Dentista
func (s *dentistaSqlStore) Update(product domain.Dentista) error {
	return nil
}

// Delete elimina un Dentista
func (s *dentistaSqlStore) Delete(id int) error {
	return nil
}

// Exists verifica si un Dentista existe
func (s *dentistaSqlStore) Exists(codeValue string) bool {
	return true
}