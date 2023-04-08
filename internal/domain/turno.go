package domain

type Turno struct {
	Paciente Paciente		`json:"paciente"`
	Dentista Dentista		`json:"dentista"`
	FechaYHora string		`json:"fechaYHora"`
	Descripción string		`json:"descripción"`
}