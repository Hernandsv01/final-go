package domain

type Paciente struct {
	DNI int				`json:"matricula"`
	Apellido string		`json:"apellido"`
	Nombre string		`json:"nombre"`
	Domicilio string 	`json:"domicilio"`
	FechaAlta string	`json:"fechaAlta"`
}