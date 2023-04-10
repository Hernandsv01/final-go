package handler

import (
	"github.com/Hernandsv01/final-go.git/internal/dentista"
	"github.com/gin-gonic/gin"
)

type dentistaHandler struct {
	s dentista.Service
}

// NewDentistaHandler crea un nuevo controller de dentista
func NewDentistaHandler(s dentista.Service) *dentistaHandler {
	return &dentistaHandler{
		s: s,
	}
}

/*
	POST: agregar dentista.
	GET: traer dentista por ID.
	PUT: actualizar dentista.
	PATCH: actualizar un dentista por alguno de sus campos.
	DELETE: eliminar dentista.
*/

func (h *dentistaHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func (h *dentistaHandler) GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func (h *dentistaHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func (h *dentistaHandler) Patch() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func (h *dentistaHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func (h *dentistaHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

