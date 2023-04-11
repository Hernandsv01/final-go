package store

import (
	"database/sql"
	"fmt"
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
func (s *dentistaSqlStore) ReadAll() ([]domain.Dentista, error) {
	var dentistasList []domain.Dentista

	rows, err := s.db.Query("SELECT * FROM dentista")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
        var matricula int
        var apellido string
        var nombre string

        err = rows.Scan(&matricula, &apellido, &nombre)
        if err != nil {
            fmt.Println(err.Error())
        }

		dentistasList = append(dentistasList, domain.Dentista{Matricula: matricula, Apellido: apellido, Nombre: nombre})
    }

	return dentistasList, nil
}

// Read devuelve un Dentista por su id
func (s *dentistaSqlStore) Read(matricula int) (domain.Dentista, error) {
	rows, err := s.db.Query("SELECT * FROM dentista WHERE matricula = " + strconv.Itoa(matricula))
	if err != nil {
		return domain.Dentista{}, err
	}
	
	var dentistaRes domain.Dentista
	if rows.Next() {

        err = rows.Scan(&dentistaRes.Matricula, &dentistaRes.Apellido, &dentistaRes.Nombre)
        if err != nil {
            fmt.Println(err.Error())
        }

		return dentistaRes, nil
    } else {
		return dentistaRes, fmt.Errorf("Dentista not found: matricula=%d", matricula)
	}
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