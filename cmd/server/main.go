package main

import (
	"database/sql"
	"fmt"

	"github.com/Hernandsv01/final-go.git/cmd/server/handler"
	"github.com/Hernandsv01/final-go.git/internal/dentista"
	"github.com/Hernandsv01/final-go.git/internal/paciente"
	"github.com/Hernandsv01/final-go.git/internal/turno"
	"github.com/Hernandsv01/final-go.git/pkg/store"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/dbfinal")
	if err != nil {
		fmt.Println("Conexión fallida")
		return
	}else{
		fmt.Println("Conexión exitosa")
	}


	r := gin.Default()

	dentistaHandler := handler.NewDentistaHandler(dentista.NewService(dentista.NewRepository(store.NewDentistaSqlStore(db))))
	pacienteHandler := handler.NewPacienteHandler(paciente.NewService(paciente.NewRepository(store.NewPacienteSqlStore(db))))
	turnoHandler := handler.NewTurnoHandler(turno.NewService(
		turno.NewRepository(store.NewTurnoSqlStore(db)),
		dentista.NewService(dentista.NewRepository(store.NewDentistaSqlStore(db))),
		paciente.NewService(paciente.NewRepository(store.NewPacienteSqlStore(db))),
		))

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	dentistas := r.Group("/dentistas")
	{
		dentistas.POST("", dentistaHandler.Create())
		dentistas.GET("", dentistaHandler.GetAll())
		dentistas.GET(":matricula", dentistaHandler.GetByMatricula())
		dentistas.PUT(":matricula", dentistaHandler.Update("put"))
		dentistas.PATCH(":matricula", dentistaHandler.Update("patch"))
		dentistas.DELETE(":matricula", dentistaHandler.Delete())
	}
	pacientes := r.Group("/pacientes")
	{
		pacientes.POST("", pacienteHandler.Create())
		pacientes.GET("", pacienteHandler.GetAll())
		pacientes.GET(":dni", pacienteHandler.GetByDni())
		pacientes.PUT(":dni", pacienteHandler.Update("put"))
		pacientes.PATCH(":dni", pacienteHandler.Update("patch"))
		pacientes.DELETE(":dni", pacienteHandler.Delete())
	}
	turnos := r.Group("/turnos")
	{
		turnos.POST("", turnoHandler.Create())
		turnos.GET("", turnoHandler.GetAll())
		turnos.GET(":id", turnoHandler.GetById())
		turnos.DELETE(":id", turnoHandler.Delete())
		turnos.PUT(":id", turnoHandler.Update("put"))
		turnos.PATCH(":id", turnoHandler.Update("patch"))
		turnos.GET("/paciente/:dni", turnoHandler.GetByDni())
	}
	
	r.Run(":8080")
}