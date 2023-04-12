package store

import "github.com/Hernandsv01/final-go.git/internal/domain"

type DentistaStoreInterface interface {
	// Read devuelve un Dentista por su id
	Read(id int) (domain.Dentista, error)
	// Readall devuelve una lista con todos los dentistas
	ReadAll() ([]domain.Dentista, error)
	// Create agrega un nuevo Dentista
	Create(d domain.Dentista) error
	// Update actualiza un Dentista
	Update(d domain.Dentista) error
	// Update actualiza un Dentista en su totalidad
	UpdateFull(d domain.Dentista) error
	// Delete elimina un Dentista
	Delete(id int) error
	// Exists verifica si un Dentista existe
	Exists(matricula int) bool
}

type PacienteStoreInterface interface {
	// Read devuelve un Paciente por su id
	Read(id int) (domain.Paciente, error)
	// Readall devuelve una lista con todos los Pacientes
	ReadAll() ([]domain.Paciente, error)
	// Create agrega un nuevo Paciente
	Create(p domain.Paciente) error
	// Update actualiza un Paciente
	Update(p domain.Paciente) error
	// Update actualiza un Dentista en su totalidad
	UpdateFull(p domain.Paciente) error
	// Delete elimina un Paciente
	Delete(id int) error
	// Exists verifica si un Paciente existe
	Exists(dni int) bool
}

type TurnoStoreInterface interface {
	// Read devuelve un Turno por su id
	Read(id int) (domain.Turno, error)
	// Readall devuelve una lista con todos los Turnos
	ReadAll() ([]domain.Turno, error)
	// Create agrega un nuevo Turno
	Create(t domain.Turno) error
	// Update actualiza un Turno
	Update(t domain.Turno) error
	// Update actualiza un Dentista en su totalidad
	UpdateFull(p domain.Turno) error
	// Delete elimina un Turno
	Delete(id int) error
	// GetByDni devuelve una lista con todos los turnos de un paciente
	GetByDni(dni int) ([]domain.Turno, error)
}