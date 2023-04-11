package store

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/Hernandsv01/final-go.git/internal/domain"
)

type pacienteSqlStore struct {
	db *sql.DB
}

func NewPacienteSqlStore(db *sql.DB) PacienteStoreInterface {
	return &pacienteSqlStore{
		db: db,
	}
}

// Create agrega un nuevo Paciente
func (s *pacienteSqlStore) Create(p domain.Paciente) error {
	st, err := s.db.Prepare("INSERT INTO paciente(dni, apellido, nombre, domicilio, fecha_alta) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = st.Exec(strconv.Itoa(p.DNI), p.Apellido, p.Nombre, p.Domicilio, p.FechaAlta)
	if err != nil {
		return err
	}

	st.Close()
	return nil
}

// Read devuelve una lista con todos los Pacientes
func (s *pacienteSqlStore) ReadAll() ([]domain.Paciente, error) {
	pacientesList := make([]domain.Paciente, 0)

	rows, err := s.db.Query("SELECT * FROM paciente")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var dni int
		var apellido string
		var nombre string
		var domicilio string
		var fechaAlta string

		err = rows.Scan(&dni, &apellido, &nombre, &domicilio, &fechaAlta)
		if err != nil {
			fmt.Println(err.Error())
		}

		pacientesList = append(pacientesList, domain.Paciente{DNI: dni, Apellido: apellido, Nombre: nombre, Domicilio: domicilio, FechaAlta: fechaAlta})
	}

	return pacientesList, nil
}

// Read devuelve un Paciente por su id
func (s *pacienteSqlStore) Read(dni int) (domain.Paciente, error) {
	rows, err := s.db.Query("SELECT * FROM paciente WHERE dni = " + strconv.Itoa(dni))
	if err != nil {
		return domain.Paciente{}, err
	}

	var pacienteRes domain.Paciente
	if rows.Next() {

		err = rows.Scan(&pacienteRes.DNI, &pacienteRes.Apellido, &pacienteRes.Nombre, &pacienteRes.Domicilio, &pacienteRes.FechaAlta)
		if err != nil {
			fmt.Println(err.Error())
		}

		return pacienteRes, nil
	} else {
		return pacienteRes, fmt.Errorf("Paciente not found: dni=%d", dni)
	}
}

// Update actualiza un Paciente en su totalidad
func (s *pacienteSqlStore) UpdateFull(d domain.Paciente) error {
	st, err := s.db.Prepare("UPDATE paciente SET apellido=?, nombre=?, domicilio=?, fecha_alta=? WHERE dni=?")
	if err != nil {
		return err
	}
	res, err := st.Exec(d.Apellido, d.Nombre, d.Domicilio, d.FechaAlta, strconv.Itoa(d.DNI))
	if err != nil {
		return err
	}
	rowsAff, _ := res.RowsAffected()
	if rowsAff == 0 {
		return fmt.Errorf("No paciente with that dni was found")
	}

	st.Close()
	return nil
}

// Update actualiza un Paciente
func (s *pacienteSqlStore) Update(d domain.Paciente) error {
	if d.Apellido == "" && d.Nombre == "" {
		return fmt.Errorf("New paciente is empty")
	}

	/*
	 Los atributos fueron puestos así porque
	 no se me ocurría forma de meterlos en el exec condicionalmente
	 sin tener que agregar 20 lineas mas
	*/
	statement := "UPDATE paciente SET"
	if d.Apellido != "" {
		statement += " apellido=\"" + d.Apellido + "\""
	}
	if d.Nombre != "" {
		statement += " nombre=\"" + d.Nombre + "\""
	}
	if d.Nombre != "" {
		statement += " domicilio=\"" + d.Domicilio + "\""
	}
	if d.Nombre != "" {
		statement += " fecha_alta=\"" + d.FechaAlta + "\""
	}
	statement += " WHERE dni=?"

	st, err := s.db.Prepare(statement)
	if err != nil {
		return err
	}
	res, err := st.Exec(strconv.Itoa(d.DNI))
	if err != nil {
		return err
	}
	rowsAff, _ := res.RowsAffected()
	if rowsAff == 0 {
		return fmt.Errorf("No paciente with that dni was found")
	}

	st.Close()
	return nil
}

// Delete elimina un Paciente
func (s *pacienteSqlStore) Delete(dni int) error {
	st, err := s.db.Prepare("DELETE FROM paciente WHERE dni=?")
	if err != nil {
		return err
	}
	res, err := st.Exec(strconv.Itoa(dni))
	if err != nil {
		return err
	}
	rowsAff, _ := res.RowsAffected()
	if rowsAff == 0 {
		return fmt.Errorf("No paciente with that dni was found")
	}

	st.Close()
	return nil
}

// Exists verifica si un Paciente existe
// func (s *pacienteSqlStore) Exists(dni string) bool {
// 	return true
// }
