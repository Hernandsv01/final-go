package store

import (
	"database/sql"
	"strconv"

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


// Create agrega un nuevo Dentista
func (s *dentistaSqlStore) Create(d domain.Dentista) error {
 	st, err := s.db.Prepare("INSERT INTO dentista(matricula, apellido, nombre) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = st.Exec(strconv.Itoa(d.Matricula), d.Apellido, d.Nombre)
	if err != nil {
		return err
	}

	st.Close()
	return nil
}

// Read devuelve un Dentista por su id
func (s *dentistaSqlStore) Read(id int) (domain.Dentista, error) {
	return domain.Dentista{}, nil
}

// Update actualiza un Dentista
func (s *dentistaSqlStore) Update(d domain.Dentista) error {
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