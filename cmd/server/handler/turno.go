package handler

import (
	"net/http"
	"strconv"

	"github.com/Hernandsv01/final-go.git/internal/domain"
	"github.com/Hernandsv01/final-go.git/internal/turno"
	"github.com/gin-gonic/gin"
)

type turnoHandler struct {
	s turno.Service
}

// NewTurnoHandler crea un nuevo controller de turno
func NewTurnoHandler(s turno.Service) *turnoHandler {
	return &turnoHandler{
		s: s,
	}
}

func (h *turnoHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var t domain.Turno

		if err := c.ShouldBindJSON(&t); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Estructura inválida de turno"})
			return
		}

		err := h.s.Create(t)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, "Turno creado exitosamente")
	}
}

func (h *turnoHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := h.s.ReadAll()
		c.JSON(http.StatusOK, res)
	}
}

func (h *turnoHandler) GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}
		res, err := h.s.Read(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)

	}
}

func (h *turnoHandler) GetByDni() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("dni"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "DNI inválido"})
			return
		}
		res, err := h.s.GetByDni(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func (h *turnoHandler) Update(functionType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var t domain.Turno
		if err := c.ShouldBindJSON(&t); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Estructura inválida de turno"})
			return
		}
		
		id := c.Param("id")
		err := h.s.Update(id, t, functionType)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, "Turno updated succesfully")
	}
}

func (h *turnoHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}
		err = h.s.Delete(idInt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, "Turno borrado exitosamente")
	}
}
