package domain
 /*
	Se usaron punteros a paciente y dentista para diferenciar en el patch(update) los dos siguientes casos:
		1) Que el cliente no mande paciente.DNI o dentista.Matricula porque quiere actualizar otra cosa (proceso continúa)
		2) Que el cliente mande ese atributos pero sean 0 (caso no válido y el resultado es un error)
 */
type Turno struct {
	Id int 					`json:"id"`
	Paciente Paciente		`json:"paciente"`
	Dentista Dentista		`json:"dentista"`
	FechaYHora string		`json:"fechaYHora"`
	Descripción string		`json:"descripcion"`
}