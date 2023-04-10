package main

import (
	"database/sql"
	"fmt"

	"github.com/Hernandsv01/final-go.git/cmd/server/handler"
	"github.com/Hernandsv01/final-go.git/internal/dentista"
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
	// pacienteHandler := handler.NewPacienteHandler(paciente.NewService(paciente.NewRepository(store.NewSqlStore(db))))
	// turnoHandler := handler.NewTurnoHandler(turno.NewService(turno.NewRepository(store.NewSqlStore(db))))

	products := r.Group("/dentistas")
	{
		// products.GET("", dentistaHandler.GetAll())
		// products.GET(":id", dentistaHandler.GetByID())
		// products.GET("/search", dentistaHandler.Search())
		// products.POST("", dentistaHandler.Post())
		products.GET("/test", dentistaHandler.Create())
	}


	
	r.Run(":8080")
}