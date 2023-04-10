package store

import "github.com/Hernandsv01/final-go.git/internal/domain"

type DentistaStoreInterface interface {
	// Read devuelve un Dentista por su id
	Read(id int) (domain.Dentista, error)
	// Read devuelve un Dentista por su id
	ReadAll() ([]domain.Dentista, error)
	// Create agrega un nuevo Dentista
	Create(product domain.Dentista) error
	// Update actualiza un Dentista
	Update(product domain.Dentista) error
	// Delete elimina un Dentista
	Delete(id int) error
	// Exists verifica si un Dentista existe
	Exists(codeValue string) bool
}

type PacienteStoreInterface interface {
	// Read devuelve un Paciente por su id
	Read(id int) (domain.Paciente, error)
	// Create agrega un nuevo Paciente
	Create(product domain.Paciente) error
	// Update actualiza un Paciente
	Update(product domain.Paciente) error
	// Delete elimina un Paciente
	Delete(id int) error
	// Exists verifica si un dentista existe
	Exists(codeValue string) bool
}

type TurnoStoreInterface interface {
	// Read devuelve un Turno por su id
	Read(id int) (domain.Turno, error)
	// Create agrega un nuevo Turno
	Create(product domain.Turno) error
	// Update actualiza un Turno
	Update(product domain.Turno) error
	// Delete elimina un Turno
	Delete(id int) error
	// Exists verifica si un Turno existe
	Exists(codeValue string) bool
}