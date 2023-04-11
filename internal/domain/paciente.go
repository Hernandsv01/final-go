package domain

type Paciente struct {
	DNI int				`json:"dni"`
	Apellido string		`json:"apellido"`
	Nombre string		`json:"nombre"`
	Domicilio string 	`json:"domicilio"`
	FechaAlta string	`json:"fecha_alta"`
}