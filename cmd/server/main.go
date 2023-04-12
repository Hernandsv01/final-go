package main

import (
	"database/sql"
	"fmt"

	"github.com/Hernandsv01/final-go.git/cmd/server/handler"
	"github.com/Hernandsv01/final-go.git/internal/dentista"
	"github.com/Hernandsv01/final-go.git/internal/paciente"
	"github.com/Hernandsv01/final-go.git/internal/turno"
	"github.com/Hernandsv01/final-go.git/pkg/middleware"
	"github.com/Hernandsv01/final-go.git/pkg/store"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	/*NO ME ALCANZA EL TIEMPO PARA DOCUMENTAR, PERDÓN*/
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Println("No se pudo cargar el archivo de variables de entorno")
		return
	}

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
		dentistas.GET("", dentistaHandler.GetAll())
		dentistas.GET(":matricula", dentistaHandler.GetByMatricula())
		dentistas.POST("", middleware.Authenticate(), dentistaHandler.Create())
		dentistas.PUT(":matricula", middleware.Authenticate(), dentistaHandler.Update("put"))
		dentistas.PATCH(":matricula", middleware.Authenticate(), dentistaHandler.Update("patch"))
		dentistas.DELETE(":matricula", middleware.Authenticate(), dentistaHandler.Delete())
	}
	pacientes := r.Group("/pacientes")
	{
		pacientes.GET("", pacienteHandler.GetAll())
		pacientes.GET(":dni", pacienteHandler.GetByDni())
		pacientes.POST("", middleware.Authenticate(), pacienteHandler.Create())
		pacientes.PUT(":dni", middleware.Authenticate(), pacienteHandler.Update("put"))
		pacientes.PATCH(":dni", middleware.Authenticate(), pacienteHandler.Update("patch"))
		pacientes.DELETE(":dni", middleware.Authenticate(), pacienteHandler.Delete())
	}
	turnos := r.Group("/turnos")
	{
		turnos.GET("", turnoHandler.GetAll())
		turnos.GET(":id", turnoHandler.GetById())
		turnos.GET("/paciente/:dni", turnoHandler.GetByDni())
		turnos.POST("", middleware.Authenticate(), turnoHandler.Create())
		turnos.DELETE(":id", middleware.Authenticate(), turnoHandler.Delete())
		turnos.PUT(":id", middleware.Authenticate(), turnoHandler.Update("put"))
		turnos.PATCH(":id", middleware.Authenticate(), turnoHandler.Update("patch"))
	}
	
	r.Run(":8080")
}