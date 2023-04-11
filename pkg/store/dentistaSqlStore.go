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

// Read devuelve una lista con todos los Dentistas
func (s *dentistaSqlStore) ReadAll() ([]domain.Dentista, error) {
	dentistasList := make([]domain.Dentista, 0)

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

// Update actualiza un Dentista en su totalidad
func (s *dentistaSqlStore) UpdateFull(d domain.Dentista) error {
	st, err := s.db.Prepare("UPDATE dentista SET apellido=?, nombre=? WHERE matricula=?")
	if err != nil {
		return err
	}
	res, err := st.Exec(d.Apellido, d.Nombre, strconv.Itoa(d.Matricula))
	if err != nil {
		return err
	}
	rowsAff, _ := res.RowsAffected()
	if rowsAff == 0 {
		return fmt.Errorf("No dentista with that matricula was found")
	}

	st.Close()
	return nil
}

// Update actualiza un Dentista
func (s *dentistaSqlStore) Update(d domain.Dentista) error {
	if d.Apellido == "" && d.Nombre == "" {
		return fmt.Errorf("New dentista is empty")
	}

	/*
	  Apellido y nombre fueron puestos así porque
	  no se me ocurría forma de meterlos en el exec condicionalmente
	  sin tener que agregar 20 lineas mas
	*/
	statement := "UPDATE dentista SET"
	if d.Apellido != "" {
		statement += " apellido=\"" + d.Apellido + "\""
	}
	if d.Nombre != "" {
		statement += " nombre=\"" + d.Nombre + "\""
	}
	statement += " WHERE matricula=?"

	st, err := s.db.Prepare(statement)
	if err != nil {
		return err
	}
	res, err := st.Exec(strconv.Itoa(d.Matricula))
	if err != nil {
		return err
	}
	rowsAff, _ := res.RowsAffected()
	if rowsAff == 0 {
		return fmt.Errorf("No dentista with that matricula was found")
	}

	st.Close()
	return nil
}

// Delete elimina un Dentista
func (s *dentistaSqlStore) Delete(matricula int) error {
	st, err := s.db.Prepare("DELETE FROM dentista WHERE matricula=?")
	if err != nil {
		return err
	}
	res, err := st.Exec(strconv.Itoa(matricula))
	if err != nil {
		return err
	}
	rowsAff, _ := res.RowsAffected()
	if rowsAff == 0 {
		return fmt.Errorf("No dentista with that matricula was found")
	}

	st.Close()
	return nil
}

// Exists verifica si un Dentista existe
// func (s *dentistaSqlStore) Exists(matricula string) bool {
// 	return true
// }
