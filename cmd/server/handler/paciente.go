package handler

import (
	"net/http"
	"strconv"

	"github.com/Hernandsv01/final-go.git/internal/paciente"
	"github.com/Hernandsv01/final-go.git/internal/domain"
	"github.com/gin-gonic/gin"
)

type pacienteHandler struct {
	s paciente.Service
}

// NewPacienteHandler crea un nuevo controller de paciente
func NewPacienteHandler(s paciente.Service) *pacienteHandler {
	return &pacienteHandler{
		s: s,
	}
}

func (h *pacienteHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var p domain.Paciente

		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		p, err := h.s.Create(p)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, p)
	}
}

func (h *pacienteHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := h.s.ReadAll()
		c.JSON(http.StatusOK, res)
	}
}

func (h *pacienteHandler) GetByDni() gin.HandlerFunc {
	return func(c *gin.Context) {
		dni, err := strconv.Atoi(c.Param("dni"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dni"})
			return
		}
		res, err := h.s.Read(dni)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "The specified dni could not be found"})
			return
		}

		c.JSON(http.StatusOK, res)

	}
}

func (h *pacienteHandler) Update(functionType string) gin.HandlerFunc {
	
	return func(c *gin.Context) {
		var p domain.Paciente
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		dni := c.Param("dni")

		err := h.s.Update(dni, p, functionType)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, "Paciente updated succesfully")
	}
}

func (h *pacienteHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		dni := c.Param("dni")
		dniInt, err := strconv.Atoi(dni)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dni"})
			return
		}
		err = h.s.Delete(dniInt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, "Paciente deleted succesfully")
	}
}