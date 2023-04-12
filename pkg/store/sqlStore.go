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
	_, err = st.Exec(d.Matricula, d.Apellido, d.Nombre)
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
			return nil, err
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
			return dentistaRes, err
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
	res, err := st.Exec(d.Apellido, d.Nombre, d.Matricula)
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

	var algoDiferente bool = false
	/*
	  Apellido y nombre fueron puestos así porque
	  no se me ocurría forma de meterlos en el exec condicionalmente
	  sin tener que agregar 20 lineas mas
	*/
	statement := "UPDATE dentista SET"
	if d.Apellido != "" {
		statement += " apellido=\"" + d.Apellido + "\""
		algoDiferente = true
	}
	if d.Nombre != "" {
		statement += " nombre=\"" + d.Nombre + "\""
		algoDiferente = true
	}
	statement += " WHERE matricula=?"

	if !algoDiferente {
		return fmt.Errorf("No hay nada que cambiar en esta update")
	}
	st, err := s.db.Prepare(statement)
	if err != nil {
		return err
	}
	res, err := st.Exec(d.Matricula)
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
	res, err := st.Exec(matricula)
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
func (s *dentistaSqlStore) Exists(matricula int) bool {
	rows, err := s.db.Query("SELECT EXISTS(SELECT 1 FROM dentista WHERE matricula=" + strconv.Itoa(matricula) + ")")
	if err != nil {
		return false
	}

	var res int
	if rows.Next() {
		err = rows.Scan(&res)
		if err != nil {
			return false
		}
		if res == 1 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

/* -------------------------------------------------------------------------- */

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
	_, err = st.Exec(p.DNI, p.Apellido, p.Nombre, p.Domicilio, p.FechaAlta)
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
			return nil, err
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
			return pacienteRes, err
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
	res, err := st.Exec(d.Apellido, d.Nombre, d.Domicilio, d.FechaAlta, d.DNI)
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

	var algoDiferente bool = false
	/*
	 Los atributos fueron puestos así porque
	 no se me ocurría forma de meterlos en el exec condicionalmente
	 sin tener que agregar 20 lineas mas
	*/
	statement := "UPDATE paciente SET"
	if d.Apellido != "" {
		statement += " apellido=\"" + d.Apellido + "\""
		algoDiferente = true
	}
	if d.Nombre != "" {
		statement += " nombre=\"" + d.Nombre + "\""
		algoDiferente = true
	}
	if d.Nombre != "" {
		statement += " domicilio=\"" + d.Domicilio + "\""
		algoDiferente = true
	}
	if d.Nombre != "" {
		statement += " fecha_alta=\"" + d.FechaAlta + "\""
		algoDiferente = true
	}
	statement += " WHERE dni=?"

	if !algoDiferente {
		return fmt.Errorf("No hay nada que cambiar en esta update")
	}

	st, err := s.db.Prepare(statement)
	if err != nil {
		return err
	}
	res, err := st.Exec(d.DNI)
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
	res, err := st.Exec(dni)
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

func (s *pacienteSqlStore) Exists(dni int) bool {
	rows, err := s.db.Query("SELECT EXISTS(SELECT 1 FROM paciente WHERE dni=" + strconv.Itoa(dni) + ")")
	if err != nil {
		return false
	}

	var res int
	if rows.Next() {
		err = rows.Scan(&res)
		if err != nil {
			return false
		}
		if res == 1 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

/* -------------------------------------------------------------------------- */

type turnoSqlStore struct {
	db *sql.DB
}

func NewTurnoSqlStore(db *sql.DB) TurnoStoreInterface {
	return &turnoSqlStore{
		db: db,
	}
}

// Create agrega un nuevo Turno
func (s *turnoSqlStore) Create(t domain.Turno) error {
	if !NewPacienteSqlStore(s.db).Exists(t.Paciente.DNI) || !NewDentistaSqlStore(s.db).Exists(t.Dentista.Matricula) {
		return fmt.Errorf("No existe ningún paciente o ningún dentista con esos datos")
	}
	st, err := s.db.Prepare("INSERT INTO turno(paciente_dni, dentista_matricula, fecha_hora, descripcion) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = st.Exec(t.Paciente.DNI, t.Dentista.Matricula, t.FechaYHora, t.Descripción)
	if err != nil {
		return err
	}

	st.Close()
	return nil
}

// Read devuelve una lista con todos los Turnos
func (s *turnoSqlStore) ReadAll() ([]domain.Turno, error) {
	turnosList := make([]domain.Turno, 0)

	rows, err := s.db.Query("SELECT t.id, t.fecha_hora, t.descripcion, p.dni, p.apellido, p.nombre, p.domicilio, p.fecha_alta, d.matricula, d.apellido, d.nombre FROM turno t INNER JOIN paciente p ON t.paciente_dni=p.dni INNER JOIN dentista d ON t.dentista_matricula=d.matricula")
	if err != nil {
		return nil, err
	}

	var t domain.Turno
	for rows.Next() {

		err = rows.Scan(&t.Id, &t.FechaYHora, &t.Descripción,
			&t.Paciente.DNI, &t.Paciente.Apellido, &t.Paciente.Nombre, &t.Paciente.Domicilio, &t.Paciente.FechaAlta,
			&t.Dentista.Matricula, &t.Dentista.Apellido, &t.Dentista.Nombre,
		)
		if err != nil {
			return nil, err
		}

		turnosList = append(turnosList, t)
	}

	return turnosList, nil
}

// Read devuelve un Turno por su id
func (s *turnoSqlStore) Read(id int) (domain.Turno, error) {
	rows, err := s.db.Query("SELECT t.id, t.fecha_hora, t.descripcion, p.dni, p.apellido, p.nombre, p.domicilio, p.fecha_alta, d.matricula, d.apellido, d.nombre FROM turno t INNER JOIN paciente p ON t.paciente_dni=p.dni INNER JOIN dentista d ON t.dentista_matricula=d.matricula WHERE t.id = " + strconv.Itoa(id))
	if err != nil {
		return domain.Turno{}, err
	}

	var t domain.Turno
	if rows.Next() {

		err = rows.Scan(&t.Id, &t.FechaYHora, &t.Descripción,
			&t.Paciente.DNI, &t.Paciente.Apellido, &t.Paciente.Nombre, &t.Paciente.Domicilio, &t.Paciente.FechaAlta,
			&t.Dentista.Matricula, &t.Dentista.Apellido, &t.Dentista.Nombre,
		)
		if err != nil {
			return t, err
		}

		return t, nil
	} else {
		return t, fmt.Errorf("Turno no encontrado: id=%d", id)
	}
}

// Update actualiza un Turno en su totalidad
func (s *turnoSqlStore) UpdateFull(t domain.Turno) error {
	if (t.Paciente.DNI != 0 && !NewPacienteSqlStore(s.db).Exists(t.Paciente.DNI)) || (t.Dentista.Matricula != 0 && !NewDentistaSqlStore(s.db).Exists(t.Dentista.Matricula)) {
		return fmt.Errorf("Hubo un error con el dni del paciente o la matricula del dentista")
	}
	st, err := s.db.Prepare("UPDATE turno SET paciente_dni=?, dentista_matricula=?, fecha_hora=?, descripcion=? WHERE id=?")
	if err != nil {
		return err
	}
	res, err := st.Exec(t.Paciente.DNI, t.Dentista.Matricula, t.FechaYHora, t.Descripción, t.Id)
	if err != nil {
		return err
	}
	rowsAff, _ := res.RowsAffected()
	if rowsAff == 0 {
		return fmt.Errorf("Turno no encontrado: id=%d", t.Id)
	}

	st.Close()
	return nil
}

// Update actualiza un Turno
func (s *turnoSqlStore) Update(t domain.Turno) error {
	if t.Paciente.DNI != 0 && !NewPacienteSqlStore(s.db).Exists(t.Paciente.DNI) {
		return fmt.Errorf("Hubo un problema con el DNI del paciente")
	}
	if t.Dentista.Matricula != 0 && !NewDentistaSqlStore(s.db).Exists(t.Dentista.Matricula) {
		return fmt.Errorf("Hubo un problema con la matricula del dentista")
	}

	var algoDiferente bool = false

	/*
	 Los atributos fueron puestos así porque
	 no se me ocurría forma de meterlos en el exec condicionalmente
	 sin tener que agregar 20 lineas mas
	*/
	statement := "UPDATE turno SET"
	if t.Paciente.DNI != 0 && NewPacienteSqlStore(s.db).Exists(t.Paciente.DNI) {
		statement += " paciente_dni=\"" + strconv.Itoa(t.Paciente.DNI) + "\""
		algoDiferente = true
	}
	if t.Dentista.Matricula != 0 && NewDentistaSqlStore(s.db).Exists(t.Dentista.Matricula) {
		statement += " dentista_matricula=" + strconv.Itoa(t.Dentista.Matricula)
		algoDiferente = true
	}
	if t.Descripción != "" {
		statement += " descripcion=\"" + t.Descripción + "\""
		algoDiferente = true
	}
	if t.FechaYHora != "" {
		statement += " fecha_hora=\"" + t.FechaYHora + "\""
		algoDiferente = true
	}
	statement += " WHERE id=?"

	if !algoDiferente {
		return fmt.Errorf("No hay nada que cambiar en esta update")
	}

	st, err := s.db.Prepare(statement)
	if err != nil {
		return err
	}
	res, err := st.Exec(t.Id)
	if err != nil {
		return err
	}
	rowsAff, _ := res.RowsAffected()
	if rowsAff == 0 {
		return fmt.Errorf("No se pudo encontrar un turno con ese id: %d", t.Id)
	}

	st.Close()
	return nil
}

// Delete elimina un Turno
func (s *turnoSqlStore) Delete(id int) error {
	st, err := s.db.Prepare("DELETE FROM turno WHERE id=?")
	if err != nil {
		return err
	}
	res, err := st.Exec(id)
	if err != nil {
		return err
	}
	rowsAff, _ := res.RowsAffected()
	if rowsAff == 0 {
		return fmt.Errorf("No fue encontrado ningún turno con ese id")
	}

	st.Close()
	return nil
}

func (s *turnoSqlStore) Exists(id int) bool {
	rows, err := s.db.Query("SELECT EXISTS(SELECT 1 FROM turno WHERE id=" + strconv.Itoa(id))
	if err != nil {
		return false
	}

	var res int
	if rows.Next() {
		err = rows.Scan(&res)
		if err != nil {
			return false
		}
		if res == 1 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (s *turnoSqlStore) GetByDni(dni int) ([]domain.Turno, error){
	if !NewPacienteSqlStore(s.db).Exists(dni) {
		return nil, fmt.Errorf("No se pudo encontrar un paciente con ese dni")
	}
	
	turnosList := make([]domain.Turno, 0)

	rows, err := s.db.Query("SELECT t.id, t.fecha_hora, t.descripcion, p.dni, p.apellido, p.nombre, p.domicilio, p.fecha_alta, d.matricula, d.apellido, d.nombre FROM turno t INNER JOIN paciente p ON t.paciente_dni=p.dni INNER JOIN dentista d ON t.dentista_matricula=d.matricula")
	if err != nil {
		return nil, err
	}

	var t domain.Turno
	for rows.Next() {

		err = rows.Scan(&t.Id, &t.FechaYHora, &t.Descripción,
			&t.Paciente.DNI, &t.Paciente.Apellido, &t.Paciente.Nombre, &t.Paciente.Domicilio, &t.Paciente.FechaAlta,
			&t.Dentista.Matricula, &t.Dentista.Apellido, &t.Dentista.Nombre,
		)
		if err != nil {
			return nil, err
		}

		turnosList = append(turnosList, t)
	}

	return turnosList, nil
}