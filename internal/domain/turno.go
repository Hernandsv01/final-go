package domain

type Turno struct {
	Id int 					`json:"id"`
	Paciente Paciente		`json:"paciente"`
	Dentista Dentista		`json:"dentista"`
	FechaYHora string		`json:"fechaYHora"`
	Descripción string		`json:"descripcion"`
}